package main

import (
	"fmt"
	"os"

	"markito/internal/config"
	"markito/internal/database"
	"markito/internal/database/migrations"

	"github.com/leapkit/core/db"

	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tools database <command>")
		fmt.Println("Available commands:")
		fmt.Println(" - migrate")
		fmt.Println(" - create")
		fmt.Println(" - drop")

		return
	}

	switch os.Args[1] {
	case "migrate":
		// Create database if it does not exist
		conn, err := database.Connection()
		if err != nil {
			fmt.Println("error connecting:", err)
			return
		}

		err = db.RunMigrations(migrations.All, conn)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println("✅ Migrations ran successfully")
	case "create":
		file, err := os.OpenFile(config.DatabaseURL, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil && os.IsExist(err) {
			fmt.Println("✅ Database already exists")
			return
		}

		defer file.Close()

		fmt.Println("✅ Database created successfully")

	case "drop":
		err := db.Drop(config.DatabaseURL)
		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println("✅ Database dropped successfully")

	default:
		fmt.Println("command not found")

		return
	}
}
