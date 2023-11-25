package helpers

import "markito/internal/app/config"

func documentLink(id string) string {
	return config.BaseURL + "/" + id
}
