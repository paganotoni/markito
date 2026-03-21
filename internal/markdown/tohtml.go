package markdown

import (
	"html/template"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// ToHTML converts markdown to HTML
func ToHTML(md string) template.HTML {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock | parser.Footnotes
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(md))

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	result := markdown.Render(doc, renderer)

	return template.HTML(result)
}

func ExtractName(md string) string {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(md))

	var name string
	ast.WalkFunc(doc, func(node ast.Node, entering bool) ast.WalkStatus {
		if !entering || name != "" {
			return ast.GoToNext
		}

		switch n := node.(type) {
		case *ast.Heading:
			name = extractText(n)
			return ast.Terminate
		case *ast.Paragraph:
			name = extractText(n)
			return ast.Terminate
		}

		return ast.GoToNext
	})

	name = strings.TrimSpace(name)
	if len(name) > 100 {
		name = name[:100]
	}

	return name
}

func extractText(node ast.Node) string {
	var text strings.Builder
	ast.WalkFunc(node, func(n ast.Node, entering bool) ast.WalkStatus {
		if !entering {
			return ast.GoToNext
		}
		if leaf, ok := n.(*ast.Text); ok {
			text.Write(leaf.Literal)
		}
		return ast.GoToNext
	})
	return text.String()
}
