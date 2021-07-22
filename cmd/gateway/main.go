package main

import (
	"github.com/tsundata/assistant/api/enum"
)

func main() {
	a, err := CreateApp(enum.Gateway)
	if err != nil {
		panic(err)
	}

	if err := a.Start(); err != nil {
		panic(err)
	}

	a.AwaitSignal()
}
