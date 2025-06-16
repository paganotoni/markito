package documents

import (
	"fmt"

	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

// Generates the copy element component to be reused.
func CopyEl(message, targetID, class string) Node {
	return Span(
		Classes{
			"cursor-pointer": true,
			class:            true,
		},

		hx.On("click", fmt.Sprintf("copy(this.dataset.target, `%s`)", message)),
		Data("target", targetID),
		Text("📋"),
	)
}
