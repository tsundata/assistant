package robot

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/bot/finance"
	"github.com/tsundata/assistant/internal/app/bot/org"
	_ "github.com/tsundata/assistant/internal/app/bot/plugin"
	"github.com/tsundata/assistant/internal/app/bot/todo"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/lexer"
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
	todo.Bot.Metadata.Identifier:    todo.Bot,
	org.Bot.Metadata.Identifier:     org.Bot,
	finance.Bot.Metadata.Identifier: finance.Bot,
}

type Robot struct{}

func NewRobot() *Robot {
	return &Robot{}
}

func (r *Robot) Help(in string) ([]string, error) {
	if strings.ToLower(in) == "help" {
		// todo
		return []string{"help...."}, nil
	}
	return []string{}, nil
}

func (r *Robot) ParseText(in string) ([]*lexer.Token, []string, []string, error) {
	tokens, err := lexer.ParseText(in)
	if err != nil {
		return nil, []string{}, []string{}, nil
	}
	if len(tokens) == 0 {
		return nil, []string{}, []string{}, nil
	}

	var objects []string
	var tags []string
	for _, item := range tokens {
		if item.Type == lexer.ObjectToken {
			objects = append(objects, item.Value)
		}
		if item.Type == lexer.TagToken {
			tags = append(tags, item.Value)
		}
	}

	return tokens, objects, tags, nil
}

func (r *Robot) Process(ctx context.Context, tokens []*lexer.Token, bots map[string]*pb.Bot) (out []string, err error) {
	// todo tags

	var input interface{} = tokens[0].Value
	var output interface{}
	for _, item := range bots { // fixme
		fmt.Println("[robot] run bot", item.Identifier)
		if b, ok := botMap[item.Identifier]; ok {
			out, err := b.Run(ctx, input)
			if err != nil {
				return nil, err
			}
			input = out
		}
	}

	// fixme
	switch v := output.(type) {
	case string:
		return []string{v}, nil
	}

	return []string{}, nil
}
