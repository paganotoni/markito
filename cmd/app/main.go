package main

import (
	// sqlite3 driver
	"fmt"
	"markito/internal"
	"net/http"

	_ "github.com/leapkit/leapkit/core/tools/envload"
	_ "github.com/mattn/go-sqlite3" // Database driver
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
