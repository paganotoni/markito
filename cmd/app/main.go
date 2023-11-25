package main

import (
	"fmt"
	"os"

	"markito/internal/app"
	"markito/internal/app/config"

	"github.com/leapkit/core/server"
)

func main() {
	s := server.New(
		"Markito",

		server.WithHost("0.0.0.0"),
		server.WithPort(config.Port),
	)

	if err := app.AddServices(s); err != nil {
		os.Exit(1)
	}

	if err := app.AddRoutes(s); err != nil {
		os.Exit(1)
	}

	if err := s.Start(); err != nil {
		fmt.Println(err)
	}
}
