// config package contains the configuration of the application,
// such as the database URL and the port the application will listen on
// other environment variables could in the future live here so that
// these can be used from other higher level packages.
package config

import (
	"github.com/leapkit/core/assets"
	"github.com/leapkit/core/envor"
	"github.com/leapkit/core/gloves"
	"github.com/paganotoni/tailo"
)

var (
	// Environment in which the application is running, this is useful
	// to determine the way we'll run the application, for example, if
	// we're running in production we might want to disable debug mode.
	Environment = envor.Get("GO_ENV", "development")

	// DatabaseURL to connect and interact with our database instance.
	DatabaseURL = envor.Get("DATABASE_URL", "./markito.db")

	// Port in which the web application listens on.
	Port    = envor.Get("PORT", "3000")
	BaseURL = envor.Get("BASE_URL", "http://localhost:3000")

	// SessionSecret is the secret used to sign the session cookies.
	SessionSecret = envor.Get("SESSION_SECRET", "fe1e3ad2-dd27-495b-8b45-7f4ab759269f")
	SessionName   = envor.Get("SESSION_NAME", "markito_session")

	// Options for tailo, this is used to setup the tailwindcss
	// watcher and builder.
	TailoOptions = []tailo.Option{
		tailo.UseInputPath("./internal/assets/application.css"),
		tailo.UseOutputPath("./public/application.css"),
		tailo.UseConfigPath("./tailwind.config.js"),
	}

	GlovesOptions = []gloves.Option{
		gloves.WithRunner(tailo.WatcherFn(TailoOptions...)),
		gloves.WithRunner(assets.Watcher("./internal/assets", "./public")),
		gloves.WatchExtension(".go", ".env", ".json", ".js", ".md"),
	}
)
