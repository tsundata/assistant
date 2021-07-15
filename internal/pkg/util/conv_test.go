package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBoolInt(t *testing.T) {
	require.Equal(t, BoolInt(true), 1)
	require.Equal(t, BoolInt(false), 0)
}
