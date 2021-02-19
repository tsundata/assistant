package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestByteToString(t *testing.T) {
	require.Equal(t, ByteToString([]byte("Test")), "Test")
}

func TestStringToByte(t *testing.T) {
	require.Equal(t, StringToByte("Test"), []byte("Test"))
}
