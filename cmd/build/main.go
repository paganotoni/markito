package main

import (
	_ "embed"
	"os"
	"os/exec"

	"markito/internal/config"

	"github.com/leapkit/core/assets"
	"github.com/paganotoni/tailo"
)

func main() {
	err := assets.Embed("./internal/assets", "./public")
	if err != nil {
		panic(err)
	}

	tailo.Build(config.TailoOptions...)

	cmd := exec.Command("go", "build")
	cmd.Args = append(
		cmd.Args,

		`--ldflags`, `-linkmode=external -extldflags="-static"`,
		`-tags`, `osusergo,netgo`,
		`-buildvcs=false`,
		"-o", "bin/app",
		"cmd/app/main.go",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
