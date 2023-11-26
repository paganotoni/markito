package markdown

import (
	"net/http"
)

// Parse the markdown content and return HTML
// to be used in the preview section.
func Parse(w http.ResponseWriter, r *http.Request) {
	html := ToHTML(r.FormValue("MarkdownContent"))

	w.Write([]byte(html))
}
