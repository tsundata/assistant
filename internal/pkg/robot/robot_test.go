package robot

import (
	"context"
	"fmt"
	_ "github.com/tsundata/assistant/internal/app/bot/plugin"
	"github.com/tsundata/assistant/internal/app/bot/todo"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"testing"
)

func TestRobot(t *testing.T) {
	b, err := bot.NewBot(todo.Metadata, todo.Setting, []string{"any", "filter", "save"})
	if err != nil {
		t.Fatal(err)
	}
	err = b.Run(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(b.Info())
}
