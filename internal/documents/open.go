package documents

import (
	"net/http"

	"markito/internal/markdown"

	. "maragu.dev/gomponents"
)

func Open(w http.ResponseWriter, r *http.Request) {
	documents := r.Context().Value("documentsService").(*service)

	doc, err := documents.Find(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var el Node
	el = documentEl(doc, markdown.ToHTML(doc.Content))
	el = Page(SavedEl(*doc), el)

	err = el.Render(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
