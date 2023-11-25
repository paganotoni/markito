package public

import (
	"embed"
	"path/filepath"

	"markito/internal/app/config"

	"github.com/leapkit/core/mdfs"
)

var (
	//go:embed *.css *.js
	files embed.FS

	// Folder is a mdfs instance that contains all the
	// files in the public folder.
	Folder = mdfs.New(
		files,
		filepath.Join("internal", "app", "public"),
		config.Environment,
	)
)
