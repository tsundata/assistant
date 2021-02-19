package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsUrl(t *testing.T) {
	require.True(t, IsUrl("https://github.com/tsundata/assistant"))
}

func TestIsIPv4(t *testing.T) {
	require.True(t, IsIPv4("127.0.0.1"))
	require.False(t, IsIPv4("172.888.2.1"))
}

func TestGeneratePassword(t *testing.T) {
	pwd := GeneratePassword(32, "lowercase|uppercase|numbers|hyphen|underline|space|specials|brackets|no_similar")
	require.Len(t, pwd, 32)
}

func BenchmarkGeneratePassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePassword(32, "lowercase|uppercase|numbers|hyphen|underline|space|specials|brackets|no_similar")
	}
}
