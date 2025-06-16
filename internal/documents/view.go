package documents

import (
	"html/template"

	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func view(doc *Document, html template.HTML) Node {
	return Div(
		Class("grid grid-cols-2 w-full h-full gap-5 max-h-[calc(100vh-130px)]"),
		Div(
			Class("relative max-h-inherit"),
			Textarea(
				Class("w-full h-full rounded-lg p-5 bg-white border-gray-400 resize-none box-border focus:ring-blue-400 focus:border-blue-400"),
				hx.Post("/parse"),
				hx.Trigger("keyup changed delay:200ms, change delay:200ms"),
				hx.Target("#htmlcontainer article"),
				ID("MarkdownContent"),
				Name("MarkdownContent"),
				Text(doc.Content),
			),

			CopyEl(
				"✅ Markdown copied to clipboard",
				"#MarkdownContent",
				"absolute top-3 right-3 p-3 w-10 h-10 flex items-center justify-center bg-gray-200/80 cursor-pointer rounded-md hover:bg-gray-300",
			),
		),

		Div(
			Class("relative max-h-inherit bg-white"),
			Div(
				Class("w-full h-full p-5 border border-gray-400 bg-gray-150 rounded-lg overflow-scroll"),
				ID("htmlcontainer"),
				Article(
					Class("prose prose-headings:mt-0 prose-headings:mb-1 prose-p:mt-0 prose-li:m-0"),
					Raw(string(html)),
				),
			),

			Script(Raw(`hljs.highlightAll("#htmlcontainer")`)),
			CopyEl(
				"✅ HTML copied to clipboard",
				"#MarkdownContent",
				"absolute top-2 right-3 p-3 w-10 h-10 flex items-center justify-center bg-gray-300/80 cursor-pointer rounded-md hover:bg-gray-400",
			),
		),
	)
}
