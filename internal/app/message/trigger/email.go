package trigger

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"regexp"
	"strings"
)

type Email struct {
	text  string
	email []string
}

func NewEmail() *Email {
	return &Email{}
}

func (t *Email) Cond(text string) bool {
	re := regexp.MustCompile(`(?m)` + utils.EmailRegex)
	ts := re.FindAllString(text, -1)

	if len(ts) == 0 {
		return false
	}

	t.text = text
	for _, item := range ts {
		t.text = strings.ReplaceAll(t.text, item, "")
		t.email = append(t.email, item)
	}

	t.email = clear(t.email)

	return true
}

func (t *Email) Handle(ctx *ctx.Context) {
	for _, email := range t.email {
		_, err := ctx.MsgClient.Send(context.Background(), &pb.MessageRequest{Text: "Email: " + email})
		if err != nil {
			return
		}
	}
}
