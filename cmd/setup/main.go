package main

import (
	"fmt"

	"markito/internal/config"
	"markito/internal/database"
	"markito/internal/database/migrations"

	"github.com/leapkit/core/db"

	"github.com/paganotoni/tailo"

	_ "github.com/lib/pq"
)

func main() {
	err := tailo.Setup()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("✅ Tailwind CSS setup successfully")

	err = db.Create(config.DatabaseURL)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("✅ Database created successfully")

	conn, err := database.Connection()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.RunMigrations(migrations.All, conn)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("✅ Migrations ran successfully")
}
