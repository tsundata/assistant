package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetLocalIP4(t *testing.T) {
	ip := GetLocalIP4()
	require.True(t, IsIPv4(ip))
}

func TestGetAvailablePort(t *testing.T) {
	port := GetAvailablePort()
	require.GreaterOrEqual(t, port, 0)
}
