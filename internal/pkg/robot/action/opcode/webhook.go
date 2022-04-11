package opcode

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

type Webhook struct{}

func NewWebhook() *Webhook {
	return &Webhook{}
}

func (o *Webhook) Type() int {
	return TypeAsync
}

func (o *Webhook) Doc() string {
	return "webhook [string] [string]?"
}

func (o *Webhook) Run(_ context.Context, _ *inside.Context, _ component.Component, _ []interface{}) (interface{}, error) {
	return nil, nil
}
