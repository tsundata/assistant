package trigger

import "github.com/tsundata/assistant/internal/app/message/trigger/ctx"

type Trigger interface {
	Cond(text string) bool
	Handle(ctx *ctx.Context)
}

func Triggers() []Trigger {
	return []Trigger{
		NewUrl(),
		NewEmail(),
		NewTag(),
		NewUser(),
	}
}
