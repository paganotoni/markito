package internal

import (
	"cmp"
	"net/http"
	"os"

	"markito/internal/documents"
	"markito/internal/health"
	"markito/internal/markdown"
	"markito/internal/system/assets"

	"go.leapkit.dev/core/db"
	"go.leapkit.dev/core/server"
)

// DB is the database connection builder function
// that will be used by the application based on the driver and
// connection string.
var DB = db.ConnectionFn(
	cmp.Or(os.Getenv("DATABASE_URL"), "markito.db"),
	db.WithDriver("sqlite3"),
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

	// Services that will be injected in the context
	r.Use(server.InCtxMiddleware("documentsService", documents.NewService(DB)))

	r.HandleFunc("GET /{$}", documents.New)
	r.HandleFunc("GET /health", health.Check(DB))
	r.HandleFunc("POST /parse", markdown.Parse)
	r.HandleFunc("POST /save", documents.Save)
	r.HandleFunc("GET /{id}", documents.Open)
	r.HandleFunc("GET /list/all", documents.All)

	r.Folder(assets.Manager.HandlerPattern(), assets.Manager)
	return r, nil
}
