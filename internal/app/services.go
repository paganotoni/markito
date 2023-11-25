package app

import (
	"markito/internal/app/database"
	"markito/internal/documents"

	"github.com/leapkit/core/server"
)

// AddServices is a function that will be called by the server
// to inject services in the context.
func AddServices(r *server.Instance) error {
	conn, err := database.Connection()
	if err != nil {
		return err
	}

	// Services that will be injected in the context
	r.Use(server.InCtxMiddleware("documentsService", documents.NewService(conn)))

	return nil
}
