package cli

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/anchore/syft/internal/version"
)

// versionCmd returns a cobra command that prints build version information.
func versionCmd() *cobra.Command {
	var outputFormat string

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the version of this tool",
		Long:  "Display version, commit, build date, Go runtime, and platform information.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			v := version.FromBuild()
			fmt := version.Format(outputFormat)
			return version.Print(os.Stdout, v, fmt)
		},
	}

	cmd.Flags().StringVarP(
		&outputFormat,
		"output", "o",
		string(version.TextFormat),
		"format to display version information (text, json, yaml)",
	)

	return cmd
}
