package documents

import (
	"fmt"

	lucide "github.com/eduardolat/gomponents-lucide"
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

		lucide.Copy(),
	)
}
