package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetSocketCount(t *testing.T) {
	require.GreaterOrEqual(t, GetSocketCount(), uint64(0))
}

func TestGetFDCount(t *testing.T) {
	require.GreaterOrEqual(t, GetFDCount(), uint64(0))
}
