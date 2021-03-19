package trigger

import (
	"fmt"
	"github.com/tsundata/assistant/internal/app/message/trigger/ctx"
)

type User struct {
	text string
}

func NewUser() *User {
	return &User{}
}

func (t *User) Cond(text string) bool {
	return true
}

func (t *User) Handle(ctx *ctx.Context) {
	fmt.Println("User handle", t.text)
}
