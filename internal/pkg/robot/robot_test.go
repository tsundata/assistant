package robot

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tsundata/assistant/api/pb"
	_ "github.com/tsundata/assistant/internal/app/chatbot/bot/plugin"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/system"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/todo"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"testing"
)

func TestRobotRunPlugin(t *testing.T) {
	_, err := system.Bot.RunPlugin(context.Background(), nil, bot.PluginValue{Value: "system info", Stack: []interface{}{}})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(system.Bot.Info())
}

func TestRobotProcessWorkflow(t *testing.T) {
	comp := component.MockComponent()
	r := NewRobot()
	msg, err := r.ProcessWorkflow(context.Background(), bot.Context{}, comp,
		[]*bot.Token{{Value: "info version ", Type: bot.StringToken}},
		map[string]*pb.Bot{
			system.Bot.Identifier: {
				Id:         1,
				Identifier: system.Bot.Identifier,
			},
			todo.Bot.Identifier: {
				Id:         2,
				Identifier: todo.Bot.Identifier,
			},
		})
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, msg, 2)
}
