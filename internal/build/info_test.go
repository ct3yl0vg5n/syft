package build

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGet_Defaults(t *testing.T) {
	info := Get()

	assert.Equal(t, "(dev)", info.Version)
	assert.Equal(t, "(none)", info.Commit)
	assert.Equal(t, "(unknown)", info.BuildDate)
	assert.Equal(t, runtime.Version(), info.GoVersion)
	assert.Equal(t, runtime.Compiler, info.Compiler)
	assert.Equal(t, runtime.GOOS+"/"+runtime.GOARCH, info.Platform)
}

func TestGet_PlatformFormat(t *testing.T) {
	info := Get()
	// Platform must follow the os/arch convention used by Go toolchain
	require.Contains(t, info.Platform, "/",
		"platform should be in the format os/arch")
}

func TestIsDevBuild_Default(t *testing.T) {
	// In test context, ldflags are not set, so this should be true.
	assert.True(t, IsDevBuild())
}

func TestInfo_Fields(t *testing.T) {
	info := Get()

	// Ensure none of the build info fields are accidentally left empty,
	// even in dev builds where defaults should still be populated.
	fields := map[string]string{
		"Version":   info.Version,
		"Commit":    info.Commit,
		"BuildDate": info.BuildDate,
		"GoVersion": info.GoVersion,
		"Compiler":  info.Compiler,
		"Platform":  info.Platform,
	}

	for name, val := range fields {
		assert.NotEmpty(t, val, "field %s should not be empty", name)
	}
}
