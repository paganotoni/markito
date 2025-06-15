package documents

import (
	"markito/internal/markdown"
	"net/http"

	"go.leapkit.dev/core/render"
)

func Open(w http.ResponseWriter, r *http.Request) {
	documents := r.Context().Value("documentsService").(*service)

	doc, err := documents.Find(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	rw := render.FromCtx(r.Context())
	rw.Set("doc", doc)
	rw.Set("content", doc.Content)
	rw.Set("html", markdown.ToHTML(doc.Content))

	err = rw.Render("documents/document.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
