package robot

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/finance"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/org"
	_ "github.com/tsundata/assistant/internal/app/chatbot/bot/plugin"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/system"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/todo"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/bot/trigger"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"strings"
)

func RegisterBot(ctx context.Context, bus event.Bus, bots ...*bot.Bot) error {
	for _, item := range bots {
		err := bus.Publish(ctx, enum.Chatbot, event.BotRegisterSubject, pb.Bot{
			Name:       item.Name,
			Identifier: item.Identifier,
			Detail:     item.Detail,
			Avatar:     item.Avatar,
			Extend:     "",
		})
		if err != nil {
			return err
		}
	}
	return nil
}

var botMap = map[string]*bot.Bot{
	system.Bot.Metadata.Identifier:  system.Bot,
	todo.Bot.Metadata.Identifier:    todo.Bot,
	org.Bot.Metadata.Identifier:     org.Bot,
	finance.Bot.Metadata.Identifier: finance.Bot,
}

type Robot struct{}

func NewRobot() *Robot {
	return &Robot{}
}

func (r *Robot) bot(identifier string) *bot.Bot {
	if b, ok := botMap[identifier]; ok {
		return b
	}
	return nil
}

func (r *Robot) Help(bots []*pb.Bot, in string) (map[int64][]pb.MsgPayload, error) {
	out := make(map[int64][]pb.MsgPayload)
	if strings.ToLower(in) == "help" && len(bots) > 0 {
		// command help
		for _, item := range bots {
			temp := strings.Builder{}
			if r.bot(item.Identifier) == nil {
				continue
			}
			temp.WriteString("--- ")
			temp.WriteString(item.Identifier)
			temp.WriteString("  ---\n")
			c := command.New(r.bot(item.Identifier).CommandRule)
			temp.WriteString(c.Help(in))
			out[item.Id] = []pb.MsgPayload{
				pb.TextMsg{Text: temp.String()},
			}
		}
	}
	return out, nil
}

//ParseText  tokens, objects, tags, commands
func (r *Robot) ParseText(in *pb.Message) ([]*bot.Token, []string, []string, []string, error) {
	tokens, err := bot.ParseText(in.GetText())
	if err != nil || len(tokens) == 0 {
		return nil, []string{}, []string{}, []string{}, nil
	}

	var objects []string
	var tags []string
	var commands []string
	for _, item := range tokens {
		if item.Type == bot.ObjectToken {
			objects = append(objects, item.Value)
		}
		if item.Type == bot.TagToken {
			tags = append(tags, item.Value)
		}
		if item.Type == bot.CommandToken {
			commands = append(commands, item.Value)
		}
	}

	return tokens, objects, tags, commands, nil
}

func (r *Robot) ProcessTrigger(ctx context.Context, comp component.Component, in *pb.Message) error {
	return trigger.Process(ctx, comp, in)
}

func (r *Robot) ProcessCommand(ctx context.Context, comp component.Component, bot *pb.Bot, commandText string) (map[int64][]pb.MsgPayload, error) {
	if r.bot(bot.Identifier) == nil {
		return nil, errors.New("error identifier")
	}
	c := command.New(r.bot(bot.Identifier).CommandRule)
	return c.ProcessCommand(ctx, comp, bot, commandText)
}

func (r *Robot) ProcessWorkflow(ctx context.Context, comp component.Component, tokens []*bot.Token, bots map[string]*pb.Bot) (map[int64][]pb.MsgPayload, error) {
	out := make(map[int64][]pb.MsgPayload)
	var input interface{} = tokens[0].Value
	var output interface{}
	for _, item := range bots {
		fmt.Println("[robot] run bot", item.Identifier)
		if b, ok := botMap[item.Identifier]; ok {
			out, err := b.Run(ctx, comp, input)
			if err != nil {
				return nil, err
			}
			input = out
		}

		switch v := output.(type) {
		case string:
			out[item.Id] = append(out[item.Id], pb.TextMsg{Text: v})
		case pb.MsgPayload:
			out[item.Id] = append(out[item.Id], v)
		}
	}

	return out, nil
}
