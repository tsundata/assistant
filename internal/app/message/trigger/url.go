package trigger

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"regexp"
	"strings"
)

type Url struct {
	text string
	url  []string
}

func NewUrl() *Url {
	return &Url{}
}

func (t *Url) Cond(text string) bool {
	re := regexp.MustCompile(`(?m)` + utils.UrlRegex)
	ts := re.FindAllString(text, -1)

	if len(ts) == 0 {
		return false
	}

	t.text = text
	for _, item := range ts {
		t.text = strings.ReplaceAll(t.text, item, "")
		t.url = append(t.url, item)
	}

	t.url = clear(t.url)

	return true
}

func (t *Url) Handle(ctx *ctx.Context) {
	for _, url := range t.url {
		_, err := ctx.MsgClient.Send(context.Background(), &pb.MessageRequest{Text: "URL: " + url})
		if err != nil {
			return
		}
	}
}
