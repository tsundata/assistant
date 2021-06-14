package task

import (
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

type EchoTask struct {
	bus    *event.Bus
	client *rpc.Client
}

func NewEchoTask(bus *event.Bus, client *rpc.Client) *EchoTask {
	return &EchoTask{bus: bus, client: client}
}

func (t *EchoTask) Echo(data string) (bool, error) {
	err := t.bus.Publish(event.SendMessageSubject, model.Message{Text: data})
	if err != nil {
		return false, err
	}
	return true, nil
}
