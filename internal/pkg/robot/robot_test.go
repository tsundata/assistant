package robot

import (
	"context"
	"fmt"
	_ "github.com/tsundata/assistant/internal/app/chatbot/bot/plugin"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/system"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"testing"
)

func TestRobot(t *testing.T) {
	_, err := system.Bot.RunPlugin(context.Background(), nil, bot.PluginValue{Value: "system info", Stack: make(map[string]interface{})})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(system.Bot.Info())
}
