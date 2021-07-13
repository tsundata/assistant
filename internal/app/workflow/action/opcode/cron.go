package opcode

import (
	"context"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
)

type Cron struct{}

func NewCron() *Cron {
	return &Cron{}
}

func (o *Cron) Type() int {
	return TypeAsync
}

func (o *Cron) Doc() string {
	return "cron [string]"
}

func (o *Cron) Run(_ context.Context, _ *inside.Component, _ []interface{}) (interface{}, error) {
	return nil, nil
}
