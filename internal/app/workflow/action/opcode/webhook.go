package opcode

import (
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
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

func (o *Webhook) Run(_ *inside.Context, _ []interface{}) (interface{}, error) {
	return nil, nil
}
