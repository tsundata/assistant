package trigger

import (
	"github.com/tsundata/assistant/internal/app/chatbot/trigger/ctx"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger/tags"
	"regexp"
	"strings"
)

type Tag struct {
	text string
	t    []tags.Tagger
}

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) Cond(text string) bool {
	re := regexp.MustCompile(`(?m)#(\w+)(\s+)`)
	ts := re.FindAllString(text, -1)

	if len(ts) == 0 {
		return false
	}

	t.text = text
	var items []string
	for _, item := range ts {
		t.text = strings.ReplaceAll(t.text, item, "")
		item = strings.TrimSpace(item)
		item = strings.ReplaceAll(item, "#", "")
		item = strings.ToLower(item)
		items = append(items, item)
	}

	items = clear(items)
	for _, item := range items {
		mt := tags.MapTagger(item)
		if mt != nil {
			t.t = append(t.t, mt)
		}
	}

	return true
}

func (t *Tag) Handle(ctx *ctx.Context) {
	for _, item := range t.t {
		item.Handle(ctx, t.text)
	}
}
