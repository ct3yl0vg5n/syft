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
		// use exit code 2 to distinguish application errors from OS-level errors (exit code 1)
		// note: some CI systems treat any non-zero exit code the same, but this helps with local debugging
		os.Exit(2)
	}
}
