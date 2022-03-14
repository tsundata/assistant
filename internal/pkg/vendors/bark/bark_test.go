package bark

import (
	"github.com/tsundata/assistant/internal/pkg/push"
	"testing"
)

func TestSendMessage(t *testing.T) {
	t.SkipNow()
	b := NewBark("", "")
	err := b.Send(push.Message{
		Title:   "title",
		Content: "content",
		Sound:   "",
	})
	if err != nil {
		t.Fatal(err)
	}
}
