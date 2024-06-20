package internal

import (
	"cmp"
	"embed"
	"markito/internal/documents"
	"markito/internal/helpers"
	"markito/internal/markdown"
	"markito/public"
	"os"

	"github.com/leapkit/core/assets"
	"github.com/leapkit/core/db"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/server"
	"github.com/leapkit/core/session"
	"github.com/paganotoni/tailo"
)

var (
	//go:embed *.html **/*.html
	tmpls embed.FS

	// Assets is the manager for the public assets
	// it allows to watch for changes and reload the assets
	// when changes are made.
	Assets = assets.NewManager(public.Files)

	// TailoOptions allow to define how to compile
	// the tailwind css files, which is the input and
	// what will be the output.
	TailoOptions = []tailo.Option{
		tailo.UseInputPath("internal/assets/application.css"),
		tailo.UseOutputPath("public/application.css"),
		tailo.UseConfigPath("tailwind.config.js"),
	}

	// DatabaseURL to connect and interact with our database instance.
	DatabaseURL = cmp.Or(os.Getenv("DATABASE_URL"), "markito.db")

	// DB is the database connection builder function
	// that will be used by the application based on the driver and
	// connection string.
	DB = db.ConnectionFn(DatabaseURL, db.WithDriver("sqlite3"))
)

func AddRoutes(r server.Router) error {
	// LeapKit Middleware
	r.Use(session.Middleware(
		cmp.Or(os.Getenv("SESSION_SECRET"), "d720c059-9664-4980-8169-1158e167ae57"),
		cmp.Or(os.Getenv("SESSION_NAME"), "markito_session"),
	))

	r.Use(render.Middleware(
		render.TemplateFS(tmpls, "internal"),

		render.WithDefaultLayout("layout.html"),
		render.WithHelpers(helpers.All),
		render.WithHelpers(render.AllHelpers),
		render.WithHelpers(map[string]any{
			"assetPath": Assets.PathFor,
		}),
	))

	r.HandleFunc("GET /{$}", documents.New)
	r.HandleFunc("POST /parse", markdown.Parse)
	r.HandleFunc("POST /save", documents.Save)
	r.HandleFunc("GET /{id}", documents.Open)
	r.HandleFunc("GET /list/all", documents.List)

	// Mounting the assets manager at the end of the routes
	// so that it can serve the public assets.
	r.HandleFunc(Assets.HandlerPattern(), Assets.HandlerFn)

	return nil
}

// AddServices is a function that will be called by the server
// to inject services in the context.
func AddServices(r server.Router) error {
	conn, err := DB()
	if err != nil {
		return err
	}

	// Services that will be injected in the context
	r.Use(server.InCtxMiddleware("documentsService", documents.NewService(conn)))

	return nil
}
