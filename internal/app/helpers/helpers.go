package helpers

import "github.com/leapkit/core/render"

// All used by the application to render the templates.
var All = render.MergeHelpers(
	render.AllHelpers,
	map[string]any{
		"urlHas":       urlHas,
		"hasPrefix":    hasPrefix,
		"urlWithParam": urlWithParams,
		"documentLink": documentLink,
	},
)
