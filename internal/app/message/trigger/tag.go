package trigger

import (
	"github.com/tsundata/assistant/internal/app/message/trigger/ctx"
	"github.com/tsundata/assistant/internal/app/message/trigger/tags"
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
	re := regexp.MustCompile(`(?m)#(\w+)(\s+)?`)
	ts := re.FindAllString(text, -1)

	if len(ts) > 0 {
		t.text = text
		for _, item := range ts {
			t.text = strings.ReplaceAll(t.text, item, "")
			item = strings.TrimSpace(item)
			item = strings.ReplaceAll(item, "#", "")
			item = strings.ToLower(item)
			mt := tags.MapTagger(item)
			if mt != nil {
				t.t = append(t.t, mt)
			}
		}
		return true
	}

	return false
}

func (t *Tag) Handle(ctx *ctx.Context) {
	for _, item := range t.t {
		item.Handle(ctx, t.text)
	}
}
