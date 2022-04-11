package robot

import (
	"context"
	"fmt"
	_ "github.com/tsundata/assistant/internal/app/chatbot/bot/plugin"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/todo"
	"testing"
)

func TestRobot(t *testing.T) {
	_, err := todo.Bot.RunPlugin(context.Background(), nil, "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(todo.Bot.Info())
}
