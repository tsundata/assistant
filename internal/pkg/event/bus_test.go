package event

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"log"
	"testing"
	"time"
)

func TestBus(t *testing.T) {
	mq, err := CreateRabbitmq(enum.Cron)
	if err != nil {
		t.Fatal(err)
	}
	b := NewRabbitmqBus(mq, nil)

	err = b.Publish(context.Background(), "test", "test", time.Now().String())
	if err != nil {
		t.Fatal(err)
	}
	err = b.Subscribe(context.Background(), "test", "test", func(msg *Msg) error {
		log.Println(msg)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
