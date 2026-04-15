package build

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetEnvironment_GoVersion(t *testing.T) {
	env := GetEnvironment()
	require.NotEmpty(t, env.GoVersion)
	assert.True(t, strings.HasPrefix(env.GoVersion, "go"),
		"GoVersion should start with 'go', got: %s", env.GoVersion)
	assert.Equal(t, runtime.Version(), env.GoVersion)
}

func TestGetEnvironment_Compiler(t *testing.T) {
	env := GetEnvironment()
	require.NotEmpty(t, env.Compiler)
	assert.Equal(t, runtime.Compiler, env.Compiler)
}

func TestGetEnvironment_Platform(t *testing.T) {
	env := GetEnvironment()
	require.NotEmpty(t, env.Platform)

	expected := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	assert.Equal(t, expected, env.Platform)
	assert.Contains(t, env.Platform, "/",
		"Platform should be in 'os/arch' format")
}

func TestEnvironment_String(t *testing.T) {
	env := Environment{
		GoVersion:  "go1.21.0",
		Compiler:   "gc",
		Platform:   "linux/amd64",
		CGOEnabled: false,
	}

	result := env.String()
	assert.Contains(t, result, "go=go1.21.0")
	assert.Contains(t, result, "compiler=gc")
	assert.Contains(t, result, "platform=linux/amd64")
	assert.Contains(t, result, "cgo=disabled")
}

func TestEnvironment_String_CGOEnabled(t *testing.T) {
	env := Environment{
		GoVersion:  "go1.21.0",
		Compiler:   "gc",
		Platform:   "linux/amd64",
		CGOEnabled: true,
	}

	result := env.String()
	assert.Contains(t, result, "cgo=enabled")
}
