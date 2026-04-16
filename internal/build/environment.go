package build

import (
	"fmt"
	"runtime"
)

// Environment captures the build and runtime environment details.
type Environment struct {
	GoVersion string `json:"goVersion" yaml:"goVersion"`
	Compiler  string `json:"compiler"  yaml:"compiler"`
	Platform  string `json:"platform"  yaml:"platform"`
	CGOEnabled bool  `json:"cgoEnabled" yaml:"cgoEnabled"`
}

// GetEnvironment returns the current build environment information.
func GetEnvironment() Environment {
	return Environment{
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		CGOEnabled: isCGOEnabled(),
	}
}

// isCGOEnabled reports whether CGO is enabled in the current build.
// This is determined at compile time via build tags.
func isCGOEnabled() bool {
	return cgoEnabled
}

// String returns a human-readable representation of the environment.
// Format: "go=<version> compiler=<compiler> platform=<os>/<arch> cgo=<enabled|disabled> numcpu=<n>"
func (e Environment) String() string {
	cgo := "disabled"
	if e.CGOEnabled {
		cgo = "enabled"
	}
	return fmt.Sprintf("go=%s compiler=%s platform=%s cgo=%s numcpu=%d",
		e.GoVersion, e.Compiler, e.Platform, cgo, runtime.NumCPU())
}
