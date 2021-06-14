package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func BenchmarkRound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Round(3.14159265, 3)
	}
}

func TestRound(t *testing.T) {
	require.Equal(t, float64(3), Round(3.14159265, 0))
	require.Equal(t, 3.1, Round(3.14159265, 1))
	require.Equal(t, -3.14, Round(-3.14159265, 2))
	require.Equal(t, 3.142, Round(3.14159265, 3))
}
