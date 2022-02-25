package main

import (
	"github.com/tsundata/assistant/api/enum"
	_ "go.uber.org/automaxprocs"
)

// CreateInitControllersFn
// @title Flow App API
// @version 1.0
// @license.name MIT
// @license.url https://github.com/tsundata/assistant/blob/main/LICENSE
// @host localhost:5000
// @BasePath /
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
