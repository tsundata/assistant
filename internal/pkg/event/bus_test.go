package event

import (
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/internal/pkg/app"
	"log"
	"testing"
	"time"
)

func TestBus(t *testing.T) {
	n, err := CreateNats(app.Cron)
	if err != nil {
		t.Fatal(err)
	}
	b := NewBus(n)

	err = b.Publish("test", time.Now().String())
	if err != nil {
		t.Fatal(err)
	}
	err = b.Subscribe("test", func(msg *nats.Msg) {
		log.Println(msg)
	})
	if err != nil {
		t.Fatal(err)
	}
}
