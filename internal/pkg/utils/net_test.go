package utils

import (
	"log"
	"testing"
)

func TestGetLocalIP4(t *testing.T) {
	ip := GetLocalIP4()
	if !IsIPv4(ip) {
		log.Println("error: GetLocalIP4")
	}
}

func TestGetAvailablePort(t *testing.T) {
	port := GetAvailablePort()
	if port <= 0 {
		log.Println("error: GetAvailablePort")
	}
}
