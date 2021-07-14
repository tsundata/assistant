package event

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/nats-io/nats.go"
	"reflect"
)

const DefaultEventQueue = "event_queue"

type Bus interface {
	Subscribe(ctx context.Context, subject Subject, fn nats.MsgHandler) error
	Publish(ctx context.Context, subject Subject, message interface{}) error
}

type NatsBus struct {
	nc *nats.Conn
}

func NewNatsBus(nc *nats.Conn) Bus {
	return &NatsBus{nc: nc}
}

func (b *NatsBus) Subscribe(_ context.Context, subject Subject, fn nats.MsgHandler) error {
	if !(reflect.TypeOf(fn).Kind() == reflect.Func) {
		return fmt.Errorf("%s is not of type reflect.Func", reflect.TypeOf(fn).Kind())
	}
	_, err := b.nc.QueueSubscribe(string(subject), DefaultEventQueue, fn)
	if err != nil {
		return err
	}
	return nil
}

func (b *NatsBus) Publish(_ context.Context, subject Subject, message interface{}) error {
	ec, err := nats.NewEncodedConn(b.nc, nats.JSON_ENCODER)
	if err != nil {
		return err
	}
	return ec.Publish(string(subject), message)
}

var ProviderSet = wire.NewSet(NewNatsBus)
