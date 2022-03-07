package trigger

import (
	"bytes"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/util"
	"regexp"
	"strings"
)

type Url struct {
	text    string
	url     []string
	message *pb.Message
}

func NewUrl() *Url {
	return &Url{}
}

func (t *Url) Cond(message *pb.Message) bool {
	t.message = message
	re := regexp.MustCompile(`(?m)` + util.UrlRegex)
	ts := re.FindAllString(message.GetText(), -1)

	if len(ts) == 0 {
		return false
	}

	t.text = message.GetText()
	for _, item := range ts {
		t.text = strings.ReplaceAll(t.text, item, "")
		t.url = append(t.url, item)
	}

	t.url = clear(t.url)

	return true
}

func (t *Url) Handle(ctx context.Context, comp component.Component) {
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
		reply, err := comp.Middle().CreatePage(ctx, &pb.PageRequest{
			Page: &pb.Page{
				Title:   title,
				Content: util.ByteToString(resp.Body()),
				Type:    "html",
			},
		})
		if err != nil {
			return
		}

		// send message
		err = comp.GetBus().Publish(ctx, enum.Message, event.MessageChannelSubject, pb.Message{
			GroupId: t.message.GetGroupId(),
			Text:    fmt.Sprintf("Archive URL: %s\nPage: %s", url, reply.GetText()),
		})
		if err != nil {
			return
		}
	}
}
