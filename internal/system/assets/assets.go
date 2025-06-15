package assets

import (
	"embed"

	"go.leapkit.dev/core/assets"
)

var (
	//go:embed *
	Files embed.FS

	// Manager is the assets manager instance
	Manager = assets.NewManager(Files, "/public")
)
