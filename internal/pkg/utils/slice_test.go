package utils

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
}
