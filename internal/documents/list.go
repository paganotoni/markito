package documents

import (
	"net/http"

	"github.com/leapkit/core/render"
)

func List(w http.ResponseWriter, r *http.Request) {
	service := r.Context().Value("documentsService").(*service)

	docs, err := service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rw := render.FromCtx(r.Context())
	rw.Set("docs", docs)

	if err = rw.Render("documents/list.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
