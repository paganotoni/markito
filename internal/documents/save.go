package documents

import (
	"fmt"
	"net/http"

	"github.com/lithammer/shortuuid/v4"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"markito/internal/markdown"
)

// Save the content of the document in the database, if there is an id present it
// will update the document, otherwise it will create a new one.
func Save(w http.ResponseWriter, r *http.Request) {
	documents := r.Context().Value("documentsService").(*service)

	content := r.FormValue("MarkdownContent")
	name := markdown.ExtractName(content)
	id := r.FormValue("DocumentID")
	if id == "" {
		id = shortuuid.New()
	}

	err := documents.Save(id, name, content)
	if err != nil {
		fmt.Println("Error saving document", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Hx-Push", "/"+id)

	el := savedEl(Document{
		ID:      id,
		Name:    name,
		Content: content,
	})

	err = el.Render(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	oobRaw := A(
		ID("raw-link"),
		Href("/raw/"+id+"/html"),
		Target("_blank"),
		Class("absolute top-2 right-14 p-3 w-10 h-10 flex items-center justify-center bg-gray-300/80 cursor-pointer rounded-md hover:bg-gray-400 text-xs font-semibold"),
		Attr("hx-swap-oob", "true"),
		Text("RAW"),
	)

	err = oobRaw.Render(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
