package event

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/api/enum"
	"log"
	"testing"
	"time"
)

func TestBus(t *testing.T) {
	n, err := CreateNats(enum.Cron)
	if err != nil {
		t.Fatal(err)
	}
	b := NewNatsBus(n, nil, nil)

	err = b.Publish(context.Background(), "test", time.Now().String())
	if err != nil {
		t.Fatal(err)
	}
	err = b.Subscribe(context.Background(), "test", func(msg *nats.Msg) {
		log.Println(msg)
	})
	if err != nil {
		t.Fatal(err)
	}
}
