package trigger

import (
	"fmt"
	"github.com/tsundata/assistant/internal/app/message/trigger/ctx"
)

type Email struct{
	text string
}

func NewEmail() *Email {
	return &Email{}
}

func (t *Email) Cond(text string) bool {
	return true
}

func (t *Email) Handle(ctx *ctx.Context) {
	fmt.Println("Email handle", t.text)
}
