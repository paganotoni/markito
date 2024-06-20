package main

import (
	"fmt"
	"markito/internal"

	"github.com/leapkit/core/tools/rebuilder"
	"github.com/paganotoni/tailo"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := rebuilder.Start(
		"cmd/app/main.go",

		// Run the tailo watcher so when changes are made to
		// the html code it rebuilds css.
		rebuilder.WithRunner(tailo.WatcherFn(internal.TailoOptions...)),

		// Run the assets watcher.
		rebuilder.WithRunner(internal.Assets.Watch),
		rebuilder.WatchExtension(".go", ".css", ".js"),
	)

	if err != nil {
		fmt.Println(err)
	}
}
