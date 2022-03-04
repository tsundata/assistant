package robot

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/bot/finance"
	"github.com/tsundata/assistant/internal/app/bot/org"
	_ "github.com/tsundata/assistant/internal/app/bot/plugin"
	"github.com/tsundata/assistant/internal/app/bot/todo"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
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

func (r *Robot) bot(identifier string) *bot.Bot {
	if b, ok := botMap[identifier]; ok {
		return b
	}
	return nil
}

func (r *Robot) Help(bots []*pb.Bot, in string) ([]string, error) {
	if strings.ToLower(in) == "help" && len(bots) > 0 {
		out := strings.Builder{}
		// command help
		for _, item := range bots {
			if r.bot(item.Identifier) == nil {
				continue
			}
			out.WriteString("--- ")
			out.WriteString(item.Identifier)
			out.WriteString("  ---\n")
			c := command.New(r.bot(item.Identifier).CommandRule)
			out.WriteString(c.Help(in))
		}

		return []string{out.String()}, nil
	}
	return []string{}, nil
}

func (r *Robot) ParseText(in string) ([]*bot.Token, []string, []string, []string, error) {
	tokens, err := bot.ParseText(in)
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

func (r *Robot) ParseCommand(ctx context.Context, comp command.Component, identifier, in string) (out []string, err error) {
	if r.bot(identifier) == nil {
		return nil, errors.New("error identifier")
	}
	c := command.New(r.bot(identifier).CommandRule)
	return c.ParseCommand(ctx, comp, in)
}

func (r *Robot) Process(ctx context.Context, tokens []*bot.Token, bots map[string]*pb.Bot) (out []string, err error) {
	// todo tags

	var input interface{} = tokens[0].Value
	var output interface{}
	for _, item := range bots {
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
