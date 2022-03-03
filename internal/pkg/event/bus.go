package event

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tsundata/assistant/internal/pkg/log"
	"go.uber.org/zap"
	"reflect"
)

type Msg struct {
	Subject Subject
	Data    []byte
}

type MsgHandler func(msg *Msg) error

type Bus interface {
	Subscribe(ctx context.Context, service string, subject Subject, fn MsgHandler) error
	Publish(ctx context.Context, service string, subject Subject, message interface{}) error
}

type NatsBus struct {
	conn   *amqp.Connection
	logger log.Logger
}

func NewNatsBus(conn *amqp.Connection, logger log.Logger) Bus {
	return &NatsBus{conn: conn, logger: logger}
}

func (b *NatsBus) Subscribe(_ context.Context, service string, subject Subject, fn MsgHandler) error {
	if !(reflect.TypeOf(fn).Kind() == reflect.Func) {
		return fmt.Errorf("%s is not of type reflect.Func", reflect.TypeOf(fn).Kind())
	}

	if b.logger != nil {
		b.logger.Info("[event] bus subscribe", zap.Any("subject", subject))
	}

	// amqp declare
	ch, err := declare(b.conn, service, subject)
	if err != nil {
		return err
	}

	err = ch.Qos(1, 0, false)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		string(subject),
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		defer func() {
			_ = ch.Close()
		}()
		for d := range msgs {
			fmt.Println("[event] received a message:", subject, string(d.Body))
			err = fn(&Msg{
				Subject: subject,
				Data:    d.Body,
			})
			if err != nil {
				fmt.Println(err)
				err = d.Reject(false)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("[event] publish dlx", service, subject, string(d.Body))
				continue
			}
			err = d.Ack(false)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("[event] received end", subject)
	}()

	return nil
}

func (b *NatsBus) Publish(_ context.Context, service string, subject Subject, message interface{}) error {
	if b.logger != nil {
		b.logger.Info("[event] bus publish", zap.Any("subject", subject), zap.Any("message", message))
	}

	// data marshal
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// amqp declare
	ch, err := declare(b.conn, service, subject)
	if err != nil {
		return err
	}
	defer func() {
		_ = ch.Close()
	}()

	return ch.Publish(
		service,
		string(subject),
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         data,
		},
	)
}

func declare(conn *amqp.Connection, service string, subject Subject) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		service,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		fmt.Sprintf("%s_dlx", subject),
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		string(subject),
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange": fmt.Sprintf("%s_dlx", subject),
		},
	)
	if err != nil {
		return nil, err
	}

	qDlx, err := ch.QueueDeclare(
		fmt.Sprintf("%s_dlx", subject),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,
		q.Name,
		service,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		qDlx.Name,
		"",
		fmt.Sprintf("%s_dlx", subject),
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

var ProviderSet = wire.NewSet(NewNatsBus)
