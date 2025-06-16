package documents

import (
	_ "embed"
	"net/http"

	"markito/internal/markdown"

	lucide "github.com/eduardolat/gomponents-lucide"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

//go:embed default.md
var defaultMarkdown string

func New(w http.ResponseWriter, r *http.Request) {
	saveBtn := A(
		Href("#"),
		hx.Target("#link-container"),
		hx.Post("/save"),
		hx.Include("[name='MarkdownContent'], [id='DocumentID']"),

		hx.On("after-request", "htmx.trigger('#flash', 'htmx:notify', {message: '✅ Document saved successfully.'})"),
		Class("bg-blue-500 hover:bg-blue-600 text-white py-2.5 px-2 rounded-md flex flex-row gap-2"),

		lucide.Save(),
		Text("Save Document"),
	)

	doc := Document{
		Content: defaultMarkdown,
	}

	el := view(&doc, markdown.ToHTML(doc.Content))

	err := Page(saveBtn, el).Render(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
