package helpers

import (
	"github.com/leapkit/core/hctx"
	"github.com/leapkit/core/render"
)

// All used by the application to render the templates.
var All = hctx.Merge(
	hctx.Map{
		"urlHas":       urlHas,
		"hasPrefix":    hasPrefix,
		"urlWithParam": urlWithParams,
		"documentLink": documentLink,
	},
	render.AllHelpers,
)
