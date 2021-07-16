package trigger

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/model"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/event"
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

func (t *User) Handle(ctx context.Context, comp *ctx.Component) {
	for _, user := range t.user {
		if comp.User == nil {
			continue
		}

		res, err := comp.User.GetUserByName(ctx, &pb.UserRequest{Name: user})
		if err != nil {
			comp.Logger.Error(err)
			continue
		}

		err = comp.Bus.Publish(ctx, event.SendMessageSubject, model.Message{
			Text: fmt.Sprintf("User: @%s\nID: %d\nMobile: %s\nRemark: %s", user, res.Id, res.Mobile, res.Remark),
		})
		if err != nil {
			return
		}
	}
}
