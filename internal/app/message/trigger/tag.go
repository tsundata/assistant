package trigger

import (
	"github.com/tsundata/assistant/internal/app/message/trigger/tags"
	"regexp"
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
		for _, item := range ts {
			mt := tags.MapTagger(item)
			if mt != nil {
				t.t = append(t.t, mt)
			}
		}
		return true
	}

	return false
}

func (t *Tag) Handle() {
	for _, item := range t.t {
		item.Handle(t.text)
	}
}
