package internal

import (
	"cmp"
	"embed"
	"markito/internal/documents"
	"markito/internal/helpers"
	"markito/internal/markdown"
	"markito/public"
	"net/http"
	"os"

	"github.com/leapkit/leapkit/core/assets"
	"github.com/leapkit/leapkit/core/db"
	"github.com/leapkit/leapkit/core/render"
	"github.com/leapkit/leapkit/core/server"
	"github.com/paganotoni/tailo"
)

var (
	//go:embed *.html **/*.html
	tmpls embed.FS

	// TailoOptions allow to define how to compile
	// the tailwind css files, which is the input and
	// what will be the output.
	TailoOptions = []tailo.Option{
		tailo.UseInputPath("internal/assets/application.css"),
		tailo.UseOutputPath("public/application.css"),
		tailo.UseConfigPath("tailwind.config.js"),
	}

	// DB is the database connection builder function
	// that will be used by the application based on the driver and
	// connection string.
	DB = db.ConnectionFn(
		cmp.Or(os.Getenv("DATABASE_URL"), "markito.db"),
		db.WithDriver("sqlite3"),
	)
)

type Server interface {
	Addr() string
	Handler() http.Handler
}

func New() (Server, error) {
	r := server.New(
		server.WithHost(cmp.Or(os.Getenv("HOST"), "0.0.0.0")),
		server.WithPort(cmp.Or(os.Getenv("PORT"), "3000")),

		server.WithSession(
			cmp.Or(os.Getenv("SESSION_SECRET"), "d720c059-9664-4980-8169-1158e167ae57"),
			cmp.Or(os.Getenv("SESSION_NAME"), "markito_session"),
		),
	)

	assets := assets.NewManager(public.Files)
	r.Use(render.Middleware(
		render.TemplateFS(tmpls, "internal"),

		render.WithDefaultLayout("layout.html"),
		render.WithHelpers(helpers.All),
		render.WithHelpers(render.AllHelpers),
		render.WithHelpers(map[string]any{
			"assetPath": assets.PathFor,
		}),
	))

	// Services that will be injected in the context
	r.Use(server.InCtxMiddleware("documentsService", documents.NewService(DB)))

	r.HandleFunc("GET /{$}", documents.New)
	r.HandleFunc("POST /parse", markdown.Parse)
	r.HandleFunc("POST /save", documents.Save)
	r.HandleFunc("GET /{id}", documents.Open)
	r.HandleFunc("GET /list/all", documents.List)

	// Mounting the assets manager at the end of the routes
	// so that it can serve the public assets.
	r.HandleFunc(assets.HandlerPattern(), assets.HandlerFn)

	return r, nil
}
