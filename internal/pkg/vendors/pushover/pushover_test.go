package pushover

import (
	"fmt"
	"testing"
)

func TestPushMessage(t *testing.T) {
	t.SkipNow()
	p := NewPushover("", "")
	resp, err := p.PushMessage(Message{
		Title:   "test",
		Message: "content",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
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
