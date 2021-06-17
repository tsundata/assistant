package util

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

func TestPing(t *testing.T) {
	require.True(t, Ping("www.example.com"))
	require.False(t, Ping("www.this-is-a-domain-name-that-does-not-exist.com"))
}
