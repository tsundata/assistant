package util

import (
	"fmt"
	"os"
	"strings"
)

func GetSocketCount() uint64 {
	var socketCount uint64
	pid := os.Getpid()
	base := fmt.Sprintf("/proc/%d/fd", pid)
	fds, err := os.ReadDir(base)
	if err != nil {
		return 0
	}
	for _, fd := range fds {
		sl, err := os.Readlink(fmt.Sprintf("%s/%s", base, fd.Name()))
		// ignore close
		if err != nil && !strings.HasSuffix(err.Error(), "no such file or directory") {
			continue
		}

		if strings.Contains(sl, "socket") {
			socketCount++
		}
	}
	return socketCount
}

func GetFDCount() uint64 {
	pid := os.Getpid()
	path := fmt.Sprintf("/proc/%d/fd", pid)
	fds, err := os.ReadDir(path)
	if err != nil {
		return 0
	}
	return uint64(len(fds))
}
