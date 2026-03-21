package documents

import (
	"net/http"

	lucide "github.com/eduardolat/gomponents-lucide"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
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
		H2(Class("bold text-lg underline mb-4"), Textf("Total (%d)", len(docs))),
		Table(
			Class("w-full border-collapse border border-gray-300 bg-white"),
			THead(
				Tr(
					Class("bg-gray-100"),
					Th(Class("border border-gray-300 p-3 text-left"), Text("Name")),
					Th(Class("border border-gray-300 p-3 text-left"), Text("ID")),
					Th(Class("border border-gray-300 p-3 text-center w-24"), Text("Actions")),
				),
			),
			TBody(
				Map(docs, func(doc Document) Node {
					name := "Untitled"
					if doc.Name.Valid && doc.Name.String != "" {
						name = doc.Name.String
					}
					return Tr(
						Class("hover:bg-gray-50"),
						Td(Class("border border-gray-300 p-3"),
							A(Href(link(doc.ID)), Class("text-blue-600 hover:underline"), Text(name)),
						),
						Td(Class("border border-gray-300 p-3 text-gray-600 text-sm font-mono"), Text(doc.ID)),
						Td(Class("border border-gray-300 p-3 text-center"),
							Button(
								Class("text-red-500 hover:text-red-700 p-1 rounded hover:bg-red-50 transition-colors"),
								hx.Delete("/list/all/"+doc.ID),
								hx.Confirm("Are you sure you want to delete this document?"),
								hx.Target("closest tr"),
								hx.Swap("outerHTML swap:300ms"),
								Title("Delete"),
								lucide.Trash2(Class("w-4 h-4")),
							),
						),
					)
				}),
			),
		),
	}

	el = Page(Text(""), el)
	err = el.Render(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	service := r.Context().Value("documentsService").(*service)

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	err := service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
