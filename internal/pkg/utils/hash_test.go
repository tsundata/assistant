package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSHA1(t *testing.T) {
	require.Equal(t, "7c4a8d09ca3762af61e59520943dc26494f8941b", SHA1("123456"))
}
