package documents

import (
	_ "embed"
	"markito/internal/markdown"
	"net/http"

	"go.leapkit.dev/core/render"
)

//go:embed default.md
var defaultMarkdown string

func New(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())
	rw.Set("content", defaultMarkdown)
	rw.Set("html", markdown.ToHTML(defaultMarkdown))

	err := rw.Render("documents/document.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
