package documents

import (
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func savedEl(doc Document) Node {
	return Span(
		Class("flex flex-row gap-2"),
		Input(
			Type("hidden"),
			ID("DocumentID"),
			Name("DocumentID"),
			Value(doc.ID),
		),
		Span(
			Class("hidden"),
			ID("documentIDLink"),
			Text(link(doc.ID)),
		),

		Span(
			A(
				Href(link(doc.ID)),
				Class("underline"),
				Text(doc.ID),
			),
		),

		CopyEl(
			"✅ Link copied to clipboard",
			"#documentIDLink",
			"",
		),

		Div(
			Class("bg-blue-500 p-1 px-2.5 rounded-lg cursor-pointer"),
			hx.Post("/save"),
			hx.Target("#link-container"),
			hx.Include("[name='MarkdownContent'], [id='DocumentID']"),
			hx.On("before-cleanup-element", "htmx.trigger('#flash', 'htmx:notify',  {'message': 'Document saved'})"),

			Text("Save"),
		),
	)
}
