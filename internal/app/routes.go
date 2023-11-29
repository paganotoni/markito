package app

import (
	"github.com/go-chi/chi/v5/middleware"

	"github.com/leapkit/core/render"
	"github.com/leapkit/core/server"
	"github.com/leapkit/core/session"

	"markito/internal"
	"markito/internal/app/config"
	"markito/internal/app/helpers"
	"markito/internal/documents"
	"markito/internal/markdown"
	"markito/public"
)

// AddRoutes mounts the routes for the application,
// it assumes that the base services have been injected
// in the creation of the server instance.
func AddRoutes(r *server.Instance) error {
	// Base middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// LeapKit Middleware
	r.Use(session.Middleware(config.SessionSecret, config.SessionName))
	r.Use(render.Middleware(
		internal.Templates,

		render.WithHelpers(helpers.All),
		render.WithDefaultLayout("layout.html"),
	))

	r.Get("/", documents.New)
	r.Post("/parse", markdown.Parse)
	r.Post("/save", documents.Save)
	r.Get("/{id}", documents.Open)

	// Public files that include anything thats on the
	// public folder. This is useful for files and assets.
	r.Folder("/public", public.Files)

	return nil
}
