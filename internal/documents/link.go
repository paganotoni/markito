package documents

import (
	"cmp"
	"os"
)

func link(id string) string {
	baseURL := cmp.Or(os.Getenv("BASE_URL"), "http://localhost:3000")
	return baseURL + "/" + id
}
