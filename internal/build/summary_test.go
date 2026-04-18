package build

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteSummary(t *testing.T) {
	info := Info{
		Version:   "v1.2.3",
		Commit:    "abc1234def5678",
		BuildDate: "2024-01-15",
		GoVersion: "go1.21.0",
		Compiler:  "gc",
		Platform:  "linux/amd64",
	}

	var buf bytes.Buffer
	err := WriteSummary(&buf, info)
	require.NoError(t, err)

	out := buf.String()
	assert.Contains(t, out, "v1.2.3")
	assert.Contains(t, out, "abc1234def5678")
	assert.Contains(t, out, "2024-01-15")
	assert.Contains(t, out, "go1.21.0")
	assert.Contains(t, out, "linux/amd64")
	assert.Contains(t, out, "syft")
}

func TestWriteSummary_ContainsLabels(t *testing.T) {
	var buf bytes.Buffer
	err := WriteSummary(&buf, Get())
	require.NoError(t, err)

	out := buf.String()
	for _, label := range []string{"Version:", "Commit:", "Build Date:", "Go Version:", "Compiler:", "Platform:"} {
		assert.True(t, strings.Contains(out, label), "expected label %q in output", label)
	}
}

func TestShortVersion_LongCommit(t *testing.T) {
	info := Info{
		Version: "v0.90.0",
		Commit:  "deadbeefcafe1234",
	}
	result := ShortVersion(info)
	assert.Equal(t, "v0.90.0 (deadbee)", result)
}

func TestShortVersion_ShortCommit(t *testing.T) {
	info := Info{
		Version: "v0.90.0",
		Commit:  "abc",
	}
	result := ShortVersion(info)
	// commits shorter than 7 chars should not be appended in parentheses
	assert.Equal(t, "v0.90.0", result)
}

func TestShortVersion_ExactlySevenChars(t *testing.T) {
	// boundary case: a commit of exactly 7 chars should not be treated as "long"
	info := Info{
		Version: "v0.90.0",
		Commit:  "deadbee",
	}
	result := ShortVersion(info)
	assert.Equal(t, "v0.90.0", result)
}

func TestShortVersion_EmptyCommit(t *testing.T) {
	// if no commit is set (e.g. local dev build), version alone should be returned
	info := Info{
		Version: "v0.90.0",
		Commit:  "",
	}
	result := ShortVersion(info)
	assert.Equal(t, "v0.90.0", result)
}
