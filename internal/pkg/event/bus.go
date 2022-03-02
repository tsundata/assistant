package event

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/nats-io/nats.go"
	"github.com/newrelic/go-agent/v3/integrations/nrnats"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"go.uber.org/zap"
	"reflect"
)

const DefaultEventQueue = "event_queue"

type Bus interface {
	Subscribe(ctx context.Context, subject Subject, fn nats.MsgHandler) error
	Publish(ctx context.Context, subject Subject, message interface{}) error
}

type NatsBus struct {
	nc     *nats.Conn
	nr     *newrelic.App
	logger log.Logger
}

func NewNatsBus(nc *nats.Conn, nr *newrelic.App, logger log.Logger) Bus {
	return &NatsBus{nc: nc, nr: nr, logger: logger}
}

func (b *NatsBus) Subscribe(_ context.Context, subject Subject, fn nats.MsgHandler) error {
	if !(reflect.TypeOf(fn).Kind() == reflect.Func) {
		return fmt.Errorf("%s is not of type reflect.Func", reflect.TypeOf(fn).Kind())
	}
	if b.nr != nil {
		fn = nrnats.SubWrapper(b.nr.Application(), fn)
	}

	if b.logger != nil {
		b.logger.Info("bus subscribe", zap.Any("subject", subject))
	}
	_, err := b.nc.QueueSubscribe(string(subject), DefaultEventQueue, fn)
	if err != nil {
		return err
	}
	return nil
}

func (b *NatsBus) Publish(_ context.Context, subject Subject, message interface{}) error {
	if b.nr != nil {
		txn := b.nr.StartTransaction(fmt.Sprintf("event/%s", subject))
		defer nrnats.StartPublishSegment(txn, b.nc, string(subject)).End()
	}

	ec, err := nats.NewEncodedConn(b.nc, nats.JSON_ENCODER)
	if err != nil {
		return err
	}
	if b.logger != nil {
		b.logger.Info("bus publish", zap.Any("subject", subject), zap.Any("message", message))
	}
	return ec.Publish(string(subject), message)
}

var ProviderSet = wire.NewSet(NewNatsBus)
