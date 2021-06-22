package trigger

import (
	"bytes"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/util"
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
	re := regexp.MustCompile(`(?m)` + util.UrlRegex)
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
		// fetch html
		r := resty.New()
		resp, err := r.R().Get(url)
		if err != nil {
			return
		}
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
		if err != nil {
			return
		}
		title := doc.Find("title").Text()

		// store
		reply, err := ctx.Middle.CreatePage(context.Background(), &pb.PageRequest{
			Title:   title,
			Content: util.ByteToString(resp.Body()),
			Type:    "html",
		})
		if err != nil {
			return
		}

		// send message
		err = ctx.Bus.Publish(event.SendMessageSubject, model.Message{Text: fmt.Sprintf("Archive URL: %s\nPage: %s", url, reply.GetText())})
		if err != nil {
			return
		}
	}
}
