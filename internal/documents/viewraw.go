package documents

import (
	"net/http"

	"markito/internal/markdown"
)

func Markdown(w http.ResponseWriter, r *http.Request) {
	documents := r.Context().Value("documentsService").(*service)

	doc, err := documents.Find(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/markdown; charset=utf-8")
	w.Write([]byte(doc.Content))
}

func RenderHTML(w http.ResponseWriter, r *http.Request) {
	documents := r.Context().Value("documentsService").(*service)

	doc, err := documents.Find(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	html := markdown.ToHTML(doc.Content)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}
