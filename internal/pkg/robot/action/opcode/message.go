package opcode

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/util"
	"reflect"
)

type Message struct{}

func NewMessage() *Message {
	return &Message{}
}

func (o *Message) Type() int {
	return TypeOp
}

func (o *Message) Doc() string {
	return "message : (any -> bool)"
}

func (o *Message) Run(ctx context.Context, inCtx *inside.Context, comp component.Component, _ []interface{}) (interface{}, error) {
	if comp.GetBus() == nil {
		return false, nil
	}
	if inCtx.Value == nil {
		return false, nil
	}

	var text string
	if str, ok := inCtx.Value.(string); ok {
		text = str
	}
	if num, ok := inCtx.Value.(int64); ok {
		text = fmt.Sprintf("%d", num)
	}
	if boolean, ok := inCtx.Value.(bool); ok {
		text = fmt.Sprintf("%v", boolean)
	}

	v := reflect.ValueOf(inCtx.Value)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
		if v.Len() == 0 {
			return false, nil
		}
		b, err := json.Marshal(inCtx.Value)
		if err != nil {
			return false, err
		}
		text = util.ByteToString(b)
	}

	if text == "" {
		return false, nil
	}

	err := comp.GetBus().Publish(ctx, enum.Message, event.MessageSendSubject, pb.Message{
		GroupId:      inCtx.Message.GetGroupId(),
		UserId:       inCtx.Message.GetUserId(),
		Sender:       inCtx.Message.GetSender(),
		SenderType:   inCtx.Message.GetSenderType(),
		Receiver:     inCtx.Message.GetReceiver(),
		ReceiverType: inCtx.Message.GetReceiverType(),
		Type:         string(enum.MessageTypeText),
		Text:         text,
	})
	if err != nil {
		return false, err
	}
	inCtx.SetValue(true)
	return true, nil
}
