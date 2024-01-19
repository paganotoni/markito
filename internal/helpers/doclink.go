package helpers

import "markito/internal/config"

func documentLink(id string) string {
	return config.BaseURL + "/" + id
}
