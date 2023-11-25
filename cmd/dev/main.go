package main

import (
	"fmt"

	"github.com/leapkit/core/gloves"
	"markito/internal/app/config"
)

func main() {
	err := gloves.Start(
		"cmd/app/main.go",

		config.GlovesOptions...,
	)

	if err != nil {
		fmt.Println(err)
	}
}
