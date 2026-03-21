package documents

import (
	"net/http"

	"markito/internal/markdown"
)

func Markdown(w http.ResponseWriter, r *http.Request) {
	documents := r.Context().Value("documentsService").(*service)

	doc, err := documents.Find(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/markdown; charset=utf-8")
	w.Write([]byte(doc.Content))
}

func RenderHTML(w http.ResponseWriter, r *http.Request) {
	documents := r.Context().Value("documentsService").(*service)

	doc, err := documents.Find(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	htmlContent := markdown.ToHTML(doc.Content)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>` + doc.ID + `</title>
	<script src="https://cdn.tailwindcss.com?plugins=typography"></script>
</head>
<body class="bg-gray-50">
	<div class="max-w-4xl mx-auto px-6 py-12">
		<a href="/` + doc.ID + `" class="inline-flex items-center gap-2 mb-8 text-sm text-gray-600 hover:text-gray-900 transition-colors">
			<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m15 18-6-6 6-6"/></svg>
			Back to editor
		</a>
		<article class="prose prose-slate max-w-none prose-headings:mt-8 prose-headings:mb-4 prose-p:my-4 prose-li:my-1 prose-pre:my-6">
		` + string(htmlContent) + `
		</article>
	</div>
</body>
</html>`))
}
