package main

import (
	"fmt"
	"os"

	"github.com/anchore/syft/cmd/syft/cli"
	"github.com/anchore/syft/internal/build"
)

func main() {
	// write a short build summary to stderr for debug purposes (only in dev builds)
	if build.IsDevBuild() {
		build.WriteSummary(os.Stderr)
	}

	if err := cli.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
