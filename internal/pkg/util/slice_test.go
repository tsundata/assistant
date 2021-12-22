package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSliceDiff(t *testing.T) {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"c", "d", "a"}
	diff := StringSliceDiff(s1, s2)
	require.False(t, len(diff) != 1 && diff[0] != "b")

	var s3 []string
	diff = StringSliceDiff(s1, s3)
	require.Len(t, diff, 3)

	var s4 []string
	diff = StringSliceDiff(s4, s1)
	require.Len(t, diff, 3)
}

func TestIn(t *testing.T) {
	require.True(t, In([]string{"a", "a"}, "a"))
	require.True(t, In([]string{"a", "b", "c"}, "a"))
	require.False(t, In([]string{"a", "b", "c"}, "d"))
	require.False(t, In([]string{}, "a"))
}
