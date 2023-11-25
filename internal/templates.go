package internal

import (
	"embed"
	"path"

	"markito/internal/app/config"

	"github.com/leapkit/core/mdfs"
)

var (
	//go:embed **/**/*.html **/*.html
	tmpls embed.FS

	// Templates is a MDFS with the templates so we can use them in the application
	// and read them from disk in development.
	Templates = mdfs.New(
		tmpls,

		path.Join("internal"),
		config.Environment,
	)
)
