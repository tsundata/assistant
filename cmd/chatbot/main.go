package main

import (
	"github.com/tsundata/assistant/api/enum"
	_ "go.uber.org/automaxprocs"
)

func main() {
	a, err := CreateApp(enum.Chatbot)
	if err != nil {
		panic(err)
	}

	if err := a.Start(); err != nil {
		panic(err)
	}

	a.AwaitSignal()
}
