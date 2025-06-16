package helpers

import (
	"cmp"
	"os"
)

var baseURL = cmp.Or(os.Getenv("BASE_URL"), "http://localhost:3000")

func DocumentLink(id string) string {
	return baseURL + "/" + id
}
