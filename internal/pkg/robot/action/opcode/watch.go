package opcode

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

type Watch struct{}

func NewWatch() *Watch {
	return &Watch{}
}

func (o *Watch) Type() int {
	return TypeAsync
}

func (o *Watch) Doc() string {
	return "watch [string] [string]"
}

func (o *Watch) Run(_ context.Context, _ *inside.Context, _ component.Component, _ []interface{}) (interface{}, error) {
	return nil, nil
}
