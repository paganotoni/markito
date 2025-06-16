package helpers

import (
	"markito/internal/system/assets"
)

func AssetPath(path string) string {
	r, err := assets.Manager.PathFor(path)
	if err != nil {
		return ""
	}

	return r
}
