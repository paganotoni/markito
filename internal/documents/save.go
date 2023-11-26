package documents

import (
	"net/http"

	"github.com/leapkit/core/render"
	"github.com/lithammer/shortuuid/v4"
)

// Save the content of the document in the database, if there is an id present it
// will update the document, otherwise it will create a new one.
func Save(w http.ResponseWriter, r *http.Request) {
	documents := r.Context().Value("documentsService").(*service)

	content := r.FormValue("MarkdownContent")
	id := r.FormValue("DocumentID")
	if id == "" {
		id = shortuuid.New()
	}

	err := documents.Save(id, content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rw := render.FromCtx(r.Context())
	rw.Set("doc", Document{
		ID:      id,
		Content: content,
	})

	// Setting the new location of the document in the header
	w.Header().Set("Hx-Push", "/"+id)

	err = rw.RenderClean("documents/saved.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
