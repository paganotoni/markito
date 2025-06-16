package helpers

import (
	"go.leapkit.dev/core/render/hctx"
)

// All used by the application to render the templates.
var All = hctx.Merge(
	hctx.Map{
		"urlHas":       urlHas,
		"hasPrefix":    hasPrefix,
		"urlWithParam": urlWithParams,
		"documentLink": DocumentLink,
	},
)
