package robot

import (
	"context"
	"fmt"
	_ "github.com/tsundata/assistant/internal/app/bot/plugin"
	"github.com/tsundata/assistant/internal/app/bot/todo"
	"testing"
)

func TestRobot(t *testing.T) {
	_, err := todo.Bot.Run(context.Background(), nil, "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(todo.Bot.Info())
}
