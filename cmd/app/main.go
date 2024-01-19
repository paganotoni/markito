package main

import (
	"fmt"
	"os"

	"markito/internal"
	"markito/internal/config"

	"github.com/leapkit/core/server"
)

func main() {
	s := server.New(
		"Markito",

		server.WithHost("0.0.0.0"),
		server.WithPort(config.Port),
	)

	if err := internal.AddServices(s); err != nil {
		os.Exit(1)
	}

	if err := internal.AddRoutes(s); err != nil {
		os.Exit(1)
	}

	if err := s.Start(); err != nil {
		fmt.Println(err)
	}
}
