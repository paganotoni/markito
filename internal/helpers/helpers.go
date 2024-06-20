package helpers

import (
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/render/hctx"
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
