package markdown

import (
	"net/http"
)

func Parse(w http.ResponseWriter, r *http.Request) {
	html := ToHTML(r.FormValue("MarkdownContent"))

	w.Write([]byte(html))
}
