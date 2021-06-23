package version

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVersion(t *testing.T) {
	Version = "v1.0.0"

	require.Equal(t, "Version: v1.0.0\nGit Commit: unknown\nBuilt: unknown\nGo version: unknown\n", Info())
}
