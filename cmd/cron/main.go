package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	c.AddFunc("@every 1m", func() {
		fmt.Println(time.Now())
	})

	for {
		c.Run()
	}
}
