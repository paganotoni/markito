package main

import (
	"fmt"
	"markito/internal"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	_ "go.leapkit.dev/core/tools/envload"
)

func main() {
	s, err := internal.New()
	if err != nil {
		fmt.Println("[error] creating server:", err)
		return
	}

	fmt.Println("Server started at", s.Addr())
	err = http.ListenAndServe(s.Addr(), s.Handler())
	if err != nil {
		fmt.Println("[error] starting app:", err)
	}
}
