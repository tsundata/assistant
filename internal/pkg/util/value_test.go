package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValue(t *testing.T) {
	v1 := Variable("test")
	s, ok := v1.String()
	require.True(t, ok)
	require.Equal(t, "test", s)

	v2 := Variable(int64(123))
	i, ok := v2.Int64()
	require.True(t, ok)
	require.Equal(t, int64(123), i)

	v3 := Variable(true)
	b, ok := v3.Bool()
	require.True(t, ok)
	require.Equal(t, true, b)
}
