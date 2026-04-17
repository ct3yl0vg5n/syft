package build

import "runtime"

// Info holds build-time metadata about the binary.
type Info struct {
	Version    string `json:"version" yaml:"version"`
	Commit     string `json:"commit" yaml:"commit"`
	BuildDate  string `json:"buildDate" yaml:"buildDate"`
	GoVersion  string `json:"goVersion" yaml:"goVersion"`
	Compiler   string `json:"compiler" yaml:"compiler"`
	Platform   string `json:"platform" yaml:"platform"`
}

// These variables are set via -ldflags at build time.
var (
	version   = "(dev)"
	commit    = "(none)"
	buildDate = "(unknown)"
)

// Get returns the current build information.
func Get() Info {
	return Info{
		Version:   version,
		Commit:    commit,
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
		Platform:  runtime.GOOS + "/" + runtime.GOARCH,
	}
}

// IsDevBuild returns true if the binary was not built with version info.
// This is useful for skipping update checks or telemetry in local dev builds.
func IsDevBuild() bool {
	return version == "(dev)"
}
