package main

import (
	"fmt"
	"log/slog"
	"markito/internal"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

func main() {
	fmt.Println("Hello, World From the Main Package!")

	conn, err := internal.DB()
	if err != nil {
		slog.Error("Error connecting to the database: ", err)
		return
	}

	slog.Info("Connection to the database established!")
	pingErr := conn.Ping()
	if pingErr != nil {
		slog.Error("Error pinging the database: ", pingErr)
		return
	}

	slog.Info("Ping to the database successful!")

	var result int
	err = conn.QueryRow("SELECT count(*) FROM documents;").Scan(&result)
	if err != nil {
		slog.Error("Error querying the database: ", err)
		return
	}

	slog.Info("Query to the database successful!", "documents", result)
}
