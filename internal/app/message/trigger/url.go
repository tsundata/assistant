package trigger

import (
	"fmt"
	"github.com/tsundata/assistant/internal/app/message/trigger/ctx"
)

type Url struct{
	text string
}

func NewUrl() *Url {
	return &Url{}
}

func (t *Url) Cond(text string) bool {
	return true
}

func (t *Url) Handle(ctx *ctx.Context) {
	fmt.Println("url handle", t.text)
}
