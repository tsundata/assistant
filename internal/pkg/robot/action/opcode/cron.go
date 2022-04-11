package opcode

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
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

func (o *Cron) Run(_ context.Context, _ *inside.Context, _ component.Component, _ []interface{}) (interface{}, error) {
	return nil, nil
}
