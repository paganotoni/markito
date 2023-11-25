package app

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/leapkit/core/render"
	"github.com/leapkit/core/server"
	"github.com/leapkit/core/session"

	"markito/internal"
	"markito/internal/app/config"
	"markito/internal/app/helpers"
	"markito/internal/app/public"
	"markito/internal/documents"
	"markito/internal/markdown"
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
	r.Use(render.Middleware(internal.Templates, render.WithHelpers(helpers.All)))

	r.Get("/", documents.New)
	r.Post("/parse", markdown.Parse)
	r.Post("/save", documents.Save)
	r.Get("/{id}", documents.Open)

	// Public files that include anything thats on the
	// public folder. This is useful for files and assets.
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.FS(public.Folder))))

	return nil
}
