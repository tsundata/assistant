package opcode

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"github.com/tsundata/assistant/internal/pkg/event"
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

func (o *Message) Run(ctx context.Context, comp *inside.Component, _ []interface{}) (interface{}, error) {
	if comp.Bus == nil {
		return false, nil
	}
	if comp.Value == nil {
		return false, nil
	}

	var text string
	if str, ok := comp.Value.(string); ok {
		text = str
	}
	if num, ok := comp.Value.(int64); ok {
		text = fmt.Sprintf("%d", num)
	}
	if boolean, ok := comp.Value.(bool); ok {
		text = fmt.Sprintf("%v", boolean)
	}

	v := reflect.ValueOf(comp.Value)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
		if v.Len() == 0 {
			return false, nil
		}
		b, err := json.Marshal(comp.Value)
		if err != nil {
			return false, err
		}
		text = util.ByteToString(b)
	}

	if text == "" {
		return false, nil
	}

	err := comp.Bus.Publish(ctx, event.SendMessageSubject, pb.Message{Text: text})
	if err != nil {
		return false, err
	}
	comp.SetValue(true)
	return true, nil
}
