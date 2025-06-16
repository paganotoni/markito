package documents

import (
	"markito/internal/system/assets"

	lucide "github.com/eduardolat/gomponents-lucide"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

// Page renders the content within the layout of the application
func Page(saved, content Node) Node {
	pathFor := func(path string) string {
		p, _ := assets.Manager.PathFor(path)
		return p
	}

	return HTML(
		Lang("en"),
		Class("h-full bg-gray-100"),
		Head(
			Meta(Charset("UTF-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
			TitleEl(Text("Markito")),

			Link(Rel("preconnect"), Href("https://fonts.googleapis.com")),
			Link(Rel("preconnect"), Href("https://fonts.gstatic.com"), CrossOrigin("")),
			Link(Rel("stylesheet"), Href("https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/default.min.css")),
			Link(Href("https://fonts.googleapis.com/css2?family=Dancing+Script:wght@500&display=swap"), Rel("stylesheet")),
			Link(Rel("stylesheet"), Href(pathFor("application.css"))),

			Script(Src("https://unpkg.com/htmx.org@1.9.10")),
			Script(Src("https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/highlight.min.js")),
			Script(Src(pathFor("application.js"))),
		),
		Body(
			Class("h-full flex flex-col min-h-full"),
			Header(
				Class("bg-blue-400 p-3 px-5 flex flex-row gap-5 items-center text-white"),
				A(
					Href("/"),
					Class("font-bold text-lg heading-wide border-r border-1 border-white rounded-xl px-3"),
					H1(Text("📓 Markito")),
				),
				Div(
					Class("flex-grow"),
					A(
						Href("https://github.com/paganotoni/markito"),
						Class("font-medium cursor-pointer border border-white rounded-full items-center bg-blue-200 p-1.5 px-3 text-gray-600 text-sm inline-flex flex-row gap-2"),
						lucide.Github(),
						Text("Take a look at Markito repo on Github."),
					),
				),
				Span(
					Class("pl-4 flex flex-row items-center"),
					Div(ID("link-container"), saved),
				),
			),
			Section(
				Class("flex-grow flex flex-col gap-4 w-full"),
				Main(
					Class("px-5 h-full py-5 flex flex-col gap-5"),
					Section(
						Class("w-full px-5 relative"),
						Div(
							ID("flash"),
							hx.On("htmx:notify", "flash(this, event.detail.message)"),
							Class("hidden z-10 absolute p-3 px-4 border border-gray-300 shadow-md rounded-lg bg-gray-200 right-[20px] -t-[20px]"),
						),
					),

					content,
				),
			),
		),
	)
}
