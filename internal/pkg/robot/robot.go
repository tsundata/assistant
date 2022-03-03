package robot

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/internal/app/bot/finance"
	"github.com/tsundata/assistant/internal/app/bot/org"
	_ "github.com/tsundata/assistant/internal/app/bot/plugin"
	"github.com/tsundata/assistant/internal/app/bot/todo"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/lexer"
	"strings"
)

var botMap = map[string]*bot.Bot{
	todo.Bot.Metadata.Identifier:    todo.Bot,
	org.Bot.Metadata.Identifier:     org.Bot,
	finance.Bot.Metadata.Identifier: finance.Bot,
}

type Robot struct {
	repo repository.ChatbotRepository
}

func NewRobot(repo repository.ChatbotRepository) *Robot {
	return &Robot{repo: repo}
}

func (r *Robot) Process(ctx context.Context, in string) (out []string, err error) {
	if strings.ToLower(in) == "help" {
		// todo
		return []string{"help...."}, nil
	}

	tokens, err := lexer.ParseText(in)
	if err != nil {
		return nil, err
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

	fmt.Println("[robot] process", in, objects, tags)

	bots, err := r.repo.GetBotsByText(ctx, objects)
	if err != nil {
		return nil, err
	}

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

	switch v := output.(type) {
	case string:
		return []string{v}, nil
	}

	return []string{}, nil
}
