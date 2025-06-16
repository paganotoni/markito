package documents

import (
	"net/http"
	"strconv"

	"markito/internal/helpers"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func All(w http.ResponseWriter, r *http.Request) {
	service := r.Context().Value("documentsService").(*service)

	docs, err := service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var el Node
	el = Group{
		H2(
			Class("bold text-lg underline"),
			Text("Total ("+strconv.Itoa(len(docs))+")"),
		),
		Ul(
			Map(docs, func(doc Document) Node {
				return Li(
					A(
						Href(
							helpers.DocumentLink(doc.ID),
						),
						Text(doc.ID),
					),
				)
			}),
		),
	}

	el = Page(Text(""), el)
	err = el.Render(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
