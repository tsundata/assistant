package listener

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/internal/pkg/event"
)

func RegisterEventHandler(bus *event.Bus) error {
	err := bus.Subscribe(event.EchoSubject, func(msg *nats.Msg) {
		fmt.Println(msg)
	})
	if err != nil {
		return err
	}

	return nil
}
