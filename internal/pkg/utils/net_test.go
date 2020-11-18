package utils

import (
	"fmt"
	"testing"
)

func TestGetLocalIP4(t *testing.T) {
	ip := GetLocalIP4()
	if !IsIPv4(ip) {
		fmt.Println("error: GetLocalIP4")
	}
}

func TestGetAvailablePort(t *testing.T) {
	port := GetAvailablePort()
	if port <= 0 {
		fmt.Println("error: GetAvailablePort")
	}
}
