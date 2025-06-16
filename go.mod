module markito

go 1.24

require (
	github.com/gomarkdown/markdown v0.0.0-20231222211730-1d6d20845b47
	github.com/lithammer/shortuuid/v4 v4.0.0
	github.com/mattn/go-sqlite3 v1.14.24
	go.leapkit.dev/core v0.1.8
	maragu.dev/gomponents v1.1.0
	maragu.dev/gomponents-htmx v0.6.1
)

require (
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/gorilla/sessions v1.3.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	go.antoniopagano.com/tailo v0.0.11 // indirect
	go.leapkit.dev/tools v0.1.6 // indirect
	golang.org/x/sys v0.21.0 // indirect
)

tool (
	go.antoniopagano.com/tailo
	go.leapkit.dev/tools/db
	go.leapkit.dev/tools/dev
)
