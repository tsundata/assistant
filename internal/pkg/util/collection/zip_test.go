package collection

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestZip1(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	r := Zip(s1, s2)
	require.Len(t, r, 5)
	require.Equal(t, s1[0], r[0].Element1)
	require.Equal(t, s2[0], r[0].Element2)
	require.Equal(t, s1[2], r[2].Element1)
	require.Equal(t, s2[2], r[2].Element2)
	require.Equal(t, s1[4], r[4].Element1)
	require.Equal(t, s2[4], r[4].Element2)
}

func TestZip2(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3, 4, 5}
	r := Zip(s1, s2)
	require.Len(t, r, 3)
	require.Equal(t, s1[0], r[0].Element1)
	require.Equal(t, s2[0], r[0].Element2)
	require.Equal(t, s1[2], r[2].Element1)
	require.Equal(t, s2[2], r[2].Element2)

	r2 := Zip(s2, s1)
	require.Len(t, r2, 3)
	require.Equal(t, s1[0], r2[0].Element1)
	require.Equal(t, s2[0], r2[0].Element2)
	require.Equal(t, s1[2], r2[2].Element1)
	require.Equal(t, s2[2], r2[2].Element2)
}

func TestZip3(t *testing.T) {
	r := Zip("arg1", "arg2")
	require.Len(t, r, 0)
}
