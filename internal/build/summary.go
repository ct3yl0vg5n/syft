package build

import (
	"fmt"
	"io"
	"strings"
)

// WriteSummary writes a human-readable summary of build info to the given writer.
func WriteSummary(w io.Writer, info Info) error {
	lines := []string{
		fmt.Sprintf("Application: syft"),
		fmt.Sprintf("Version:     %s", info.Version),
		fmt.Sprintf("Commit:      %s", info.Commit),
		fmt.Sprintf("Build Date:  %s", info.BuildDate),
		fmt.Sprintf("Go Version:  %s", info.GoVersion),
		fmt.Sprintf("Compiler:    %s", info.Compiler),
		fmt.Sprintf("Platform:    %s", info.Platform),
	}

	_, err := fmt.Fprintln(w, strings.Join(lines, "\n"))
	return err
}

// ShortVersion returns a concise version string suitable for logging.
func ShortVersion(info Info) string {
	if len(info.Commit) >= 7 {
		return fmt.Sprintf("%s (%s)", info.Version, info.Commit[:7])
	}
	return info.Version
}
