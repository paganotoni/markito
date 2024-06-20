package helpers

import (
	"cmp"
	"os"
)

var baseURL = cmp.Or(os.Getenv("BASE_URL"), "http://localhost:3000")

func documentLink(id string) string {
	return baseURL + "/" + id
}
