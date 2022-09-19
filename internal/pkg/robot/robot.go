package robot

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	_ "github.com/tsundata/assistant/internal/app/bot/plugin"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/bot/trigger"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"strings"
)

type Robot struct{}

func NewRobot() *Robot {
	return &Robot{}
}

func (r *Robot) bot(identifier string) *bot.Bot {
	if b, ok := BotMap[identifier]; ok {
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
func (r *Robot) ParseText(in *pb.Message) ([]*bot.Token, []string, []string, []string, []string, error) {
	tokens, err := bot.ParseText(in.GetText())
	if err != nil || len(tokens) == 0 {
		return nil, []string{}, []string{}, []string{}, []string{}, nil
	}

	var objects []string
	var tags []string
	var commands []string
	var messages []string
	for _, item := range tokens {
		if item.Type == bot.ObjectToken {
			objects = append(objects, item.Value)
		}
		if item.Type == bot.TagToken {
			tags = append(tags, item.Value)
		}
		if item.Type == bot.MessageToken {
			messages = append(messages, item.Value)
		}
		if item.Type == bot.CommandToken {
			commands = append(commands, item.Value)
		}
	}

	return tokens, objects, tags, messages, commands, nil
}

func (r *Robot) ProcessTrigger(ctx context.Context, botCtx bot.Context, comp component.Component, in *pb.Message) error {
	return trigger.Process(ctx, botCtx, comp, in)
}

func (r *Robot) ProcessWorkflow(ctx context.Context, botCtx bot.Context, comp component.Component, tokens []*bot.Token, bots map[string]*pb.Bot) (map[int64][]pb.MsgPayload, error) {
	if len(tokens) == 0 {
		return map[int64][]pb.MsgPayload{}, nil
	}

	// sentence
	var str []string
	for _, token := range tokens {
		if token.Type == bot.StringToken {
			str = append(str, token.Value)
		}
	}
	sentence := strings.Join(str, " ")

	var err error
	result := make(map[int64][]pb.MsgPayload)
	botCtx.Input = bot.PluginValue{Value: sentence, Stack: []interface{}{}}
	fmt.Println("[robot] run workflow:", bots, botCtx.Input)
	for _, item := range bots {
		fmt.Println("[robot] 	run bot:", item.Identifier)
		if b, ok := BotMap[item.Identifier]; ok {
			botCtx.Input, err = b.RunPlugin(ctx, comp, botCtx.Input)
			if err != nil {
				return nil, err
			}

			runResult := b.WorkflowRule.RunFunc(ctx, botCtx, comp)
			fmt.Println("[robot] 		run func:", botCtx.Input, runResult)
			botCtx.Input = botCtx.Output
			botCtx.Input.Stack = []interface{}{}
			botCtx.Output = bot.PluginValue{}
			result[item.Id] = append(result[item.Id], runResult...)
		}
	}

	return result, nil
}

func (r *Robot) ProcessCommand(ctx context.Context, _ bot.Context, comp component.Component, identifier, commandText string) ([]pb.MsgPayload, error) {
	b, ok := BotMap[identifier]
	if !ok {
		return []pb.MsgPayload{}, nil
	}

	c := command.New(b.CommandRule)
	return c.ProcessCommand(ctx, comp, commandText)
}

func (r *Robot) ProcessAction(ctx context.Context, botCtx bot.Context, comp component.Component, identifier, id, value string) ([]pb.MsgPayload, error) {
	b, ok := BotMap[identifier]
	if !ok {
		return []pb.MsgPayload{}, nil
	}

	for _, rule := range b.ActionRule {
		if rule.ID == id {
			if f, ok := rule.OptionFunc[value]; ok {
				result := f(ctx, botCtx, comp)
				return result, nil
			}
		}
	}

	return []pb.MsgPayload{}, nil
}

func (r *Robot) ProcessForm(ctx context.Context, botCtx bot.Context, comp component.Component, identifier, id string) ([]pb.MsgPayload, error) {
	b, ok := BotMap[identifier]
	if !ok {
		return []pb.MsgPayload{}, nil
	}

	for _, rule := range b.FormRule {
		if rule.ID == id {
			result := rule.SubmitFunc(ctx, botCtx, comp)
			return result, nil
		}
	}

	return []pb.MsgPayload{}, nil
}

func (r *Robot) ProcessTag(ctx context.Context, botCtx bot.Context, comp component.Component, identifier, tag string) ([]pb.MsgPayload, error) {
	b, ok := BotMap[identifier]
	if !ok {
		return []pb.MsgPayload{}, nil
	}

	for _, rule := range b.TagRule {
		if rule.Tag == tag {
			result := rule.TriggerFunc(ctx, botCtx, comp)
			return result, nil
		}
	}

	return []pb.MsgPayload{}, nil
}

func RegisterBot(ctx context.Context, bus event.Bus, comp component.Component, bots ...*bot.Bot) error {
	for _, item := range bots {
		// info register
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
		// event handler
		if item.EventHandler != nil {
			err = item.EventHandler(comp)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
