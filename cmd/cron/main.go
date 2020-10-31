package main

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	_, err := c.AddFunc("@every 1m", func() {
		log.Println(time.Now())
	})
	if err != nil {
		panic(err)
	}

	for {
		c.Run()
	}
}
