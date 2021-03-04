package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestByteToString(t *testing.T) {
	require.Equal(t, "Test", ByteToString([]byte("Test")))
}

func TestStringToByte(t *testing.T) {
	require.Equal(t, []byte("Test"), StringToByte("Test"))
}
