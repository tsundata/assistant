package main

import "github.com/tsundata/assistant/internal/pkg/app"

func main() {
	a, err := CreateApp(app.Gateway)
	if err != nil {
		panic(err)
	}

	if err := a.Start(); err != nil {
		panic(err)
	}

	a.AwaitSignal()
}
