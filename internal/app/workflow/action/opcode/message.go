package opcode

import (
	"encoding/json"
	"fmt"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/model"
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

func (o *Message) Run(ctx *inside.Context, _ []interface{}) (interface{}, error) {
	if ctx.Bus == nil {
		return false, nil
	}
	if ctx.Value == nil {
		return false, nil
	}

	var text string
	if str, ok := ctx.Value.(string); ok {
		text = str
	}
	if num, ok := ctx.Value.(int64); ok {
		text = fmt.Sprintf("%d", num)
	}
	if boolean, ok := ctx.Value.(bool); ok {
		text = fmt.Sprintf("%v", boolean)
	}

	v := reflect.ValueOf(ctx.Value)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
		if v.Len() == 0 {
			return false, nil
		}
		b, err := json.Marshal(ctx.Value)
		if err != nil {
			return false, err
		}
		text = util.ByteToString(b)
	}

	if text == "" {
		return false, nil
	}

	err := ctx.Bus.Publish(event.SendMessageSubject, model.Message{Text: text})
	if err != nil {
		return false, err
	}
	ctx.SetValue(true)
	return true, nil
}
