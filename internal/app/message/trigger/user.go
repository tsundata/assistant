package trigger

import (
	"fmt"
	"github.com/tsundata/assistant/internal/app/message/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/model"
	"regexp"
	"strings"
)

type User struct {
	text string
	user []string
}

func NewUser() *User {
	return &User{}
}

func (t *User) Cond(text string) bool {
	re := regexp.MustCompile(`(?m)@(\w+)(\s+)`)
	ts := re.FindAllString(text, -1)

	if len(ts) == 0 {
		return false
	}

	t.text = text
	for _, item := range ts {
		t.text = strings.ReplaceAll(t.text, item, "")
		item = strings.TrimSpace(item)
		item = strings.ReplaceAll(item, "@", "")
		item = strings.ToLower(item)
		t.user = append(t.user, item)
	}

	t.user = clear(t.user)

	return true
}

func (t *User) Handle(ctx *ctx.Context) {
	for _, user := range t.user {
		err := ctx.Bus.Publish(event.SendMessageSubject, model.Message{Text: fmt.Sprintf("User: @%s", user)})
		if err != nil {
			return
		}
	}
}
