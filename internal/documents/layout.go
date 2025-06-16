package documents

import (
	"markito/internal/helpers"

	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func Page(saved, content Node) Node {
	return HTML(
		Lang("en"),
		Class("h-full bg-gray-100"),
		Head(
			Meta(
				Charset("UTF-8"),
			),
			Meta(
				Name("viewport"),
				Content("width=device-width, initial-scale=1.0"),
			),
			TitleEl(
				Text("Markito"),
			),
			Link(
				Rel("stylesheet"),
				Href(helpers.AssetPath("application.css")),
			),
			Script(
				Src("https://unpkg.com/hyperscript.org@0.9.8"),
			),
			Script(
				Src("https://unpkg.com/htmx.org@1.9.10"),
			),
			Link(
				Rel("preconnect"),
				Href("https://fonts.googleapis.com"),
			),
			Link(
				Rel("preconnect"),
				Href("https://fonts.gstatic.com"),
				CrossOrigin(""),
			),
			Link(
				Href("https://fonts.googleapis.com/css2?family=Dancing+Script:wght@500&display=swap"),
				Rel("stylesheet"),
			),
			Link(
				Rel("stylesheet"),
				Href("https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/default.min.css"),
			),
			Script(
				Src("https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/highlight.min.js"),
			),
			Script(
				Src(helpers.AssetPath("application.js")),
			),
		),
		Body(
			Class("h-full flex flex-col min-h-full"),
			Header(
				Class("bg-blue-400 p-3 px-5 flex flex-row gap-5 items-center text-white"),
				A(
					Href("/"),
					Class("font-bold text-lg heading-wide border-r border-1 border-white rounded-xl px-3"),
					H1(
						Text("📓 Markito"),
					),
				),
				Div(
					Class("flex-grow"),
					A(
						Href("https://github.com/paganotoni/markito"),
						Class("font-medium cursor-pointer border border-white rounded-full items-center bg-blue-200 p-1.5 px-3 text-gray-600 text-sm inline-flex flex-row gap-2"),
						SVG(
							Attr("xmlns", "http://www.w3.org/2000/svg"),
							Attr("viewBox", "0 0 30 30"),
							Width("24px"),
							Height("24px"),
							El("path",
								Attr("d", "M15,3C8.373,3,3,8.373,3,15c0,5.623,3.872,10.328,9.092,11.63C12.036,26.468,12,26.28,12,26.047v-2.051 c-0.487,0-1.303,0-1.508,0c-0.821,0-1.551-0.353-1.905-1.009c-0.393-0.729-0.461-1.844-1.435-2.526 c-0.289-0.227-0.069-0.486,0.264-0.451c0.615,0.174,1.125,0.596,1.605,1.222c0.478,0.627,0.703,0.769,1.596,0.769 c0.433,0,1.081-0.025,1.691-0.121c0.328-0.833,0.895-1.6,1.588-1.962c-3.996-0.411-5.903-2.399-5.903-5.098 c0-1.162,0.495-2.286,1.336-3.233C9.053,10.647,8.706,8.73,9.435,8c1.798,0,2.885,1.166,3.146,1.481C13.477,9.174,14.461,9,15.495,9 c1.036,0,2.024,0.174,2.922,0.483C18.675,9.17,19.763,8,21.565,8c0.732,0.731,0.381,2.656,0.102,3.594 c0.836,0.945,1.328,2.066,1.328,3.226c0,2.697-1.904,4.684-5.894,5.097C18.199,20.49,19,22.1,19,23.313v2.734 c0,0.104-0.023,0.179-0.035,0.268C23.641,24.676,27,20.236,27,15C27,8.373,21.627,3,15,3z"),
							),
						),
						Text("Take a look at Markito repo on Github."),
					),
				),
				Span(
					Class("pl-4 flex flex-row items-center"),
					Div(
						ID("link-container"),
						saved,
					),
				),
			),
			Section(
				Class("flex-grow flex flex-col gap-4 w-full"),
				Main(
					Class("px-5 h-full py-5 flex flex-col gap-5"),
					Attr("x-data"),
					Flash(),
					content,
				),
			),
		),
	)
}

func Flash() Node {
	return Section(
		Class("w-full px-5 relative"),
		Div(
			ID("flash"),
			hx.On("htmx:notify", "flash(this, event.detail.message)"),
			Class("hidden z-10 absolute p-3 px-4 border border-gray-300 shadow-md rounded-lg bg-gray-200 right-[20px] -t-[20px]"),
		),
	)
}
