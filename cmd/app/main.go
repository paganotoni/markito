package main

import (
	"cmp"
	"fmt"
	"net/http"
	"os"

	"markito/internal"

	"github.com/leapkit/core/server"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	s := server.New(
		server.WithHost(cmp.Or(os.Getenv("HOST"), "0.0.0.0")),
		server.WithPort(cmp.Or(os.Getenv("PORT"), "3000")),
	)

	if err := internal.AddServices(s); err != nil {
		os.Exit(1)
	}

	if err := internal.AddRoutes(s); err != nil {
		os.Exit(1)
	}

	fmt.Println("Server started at", s.Addr())
	err := http.ListenAndServe(s.Addr(), s.Handler())
	if err != nil {
		fmt.Println("[error] starting app:", err)
	}
}
