package pushover

import (
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/push"
	"testing"
)

func TestSendMessage(t *testing.T) {
	t.SkipNow()
	p := NewPushover("", "")
	err := p.Send(push.Message{
		Title:   "test",
		Content: "content",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestLimitations(t *testing.T) {
	t.SkipNow()
	p := NewPushover("", "")
	resp, err := p.Limitations()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
