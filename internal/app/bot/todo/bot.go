package todo

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
)

var metadata = bot.Metadata{
	Name:       "Todo",
	Identifier: enum.TodoBot,
	Detail:     "",
	Avatar:     "",
}

var setting = []bot.SettingField{
	{
		Key:      "report",
		Type:     bot.SettingItemTypeBool,
		Required: false,
		Value:    false,
	},
	{
		Key:      "time",
		Type:     bot.SettingItemTypeString,
		Required: true,
		Value:    "",
	},
}

var pluginRules = []bot.PluginRule{
	{
		Name: "any",
	},
	{
		Name:  "filter",
		Param: []interface{}{1},
	},
	{
		Name: "save",
	},
}

var commandRules = []command.Rule{
	{
		Define: `todo list`,
		Help:   `List todo`,
		Parse: func(ctx context.Context, comp command.Component, tokens []*command.Token) []string {
			//if comp.Todo() == nil {
			//	return []string{"empty client"}
			//}
			//
			//reply, err := comp.Todo().GetTodos(ctx, &pb.TodoRequest{})
			//if err != nil {
			//	return []string{"error call: " + err.Error()}
			//}
			//
			//tableString := &strings.Builder{}
			//if len(reply.Todos) > 0 {
			//	table := tablewriter.NewWriter(tableString)
			//	table.SetBorder(false)
			//	table.SetHeader([]string{"Id", "Priority", "Content", "Complete"})
			//	for _, v := range reply.Todos {
			//		table.Append([]string{strconv.Itoa(int(v.Id)), strconv.Itoa(int(v.Priority)), v.Content, util.BoolToString(v.Complete)})
			//	}
			//	table.Render()
			//}
			//if tableString.String() == "" {
			//	return []string{"Empty"}
			//}
			//
			//return []string{tableString.String()}
			return []string{}
		},
	},
	{
		Define: `todo [string]`,
		Help:   "Todo something",
		Parse: func(ctx context.Context, comp command.Component, tokens []*command.Token) []string {
			//if comp.Todo() == nil {
			//	return []string{"empty client"}
			//}
			//if len(tokens) != 2 {
			//	return []string{"error args"}
			//}
			//reply, err := comp.Todo().CreateTodo(ctx, &pb.TodoRequest{
			//	Todo: &pb.Todo{Content: tokens[1].Value},
			//})
			//if err != nil {
			//	return []string{"error call: " + err.Error()}
			//}
			//if !reply.GetState() {
			//	return []string{"failed"}
			//}
			return []string{"success"}
		},
	},
	{
		Define: `remind [string] [string]`,
		Help:   `Remind something`,
		Parse: func(ctx context.Context, comp command.Component, tokens []*command.Token) []string {
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			arg1 := tokens[1].Value
			arg2 := tokens[2].Value
			fmt.Println(arg1, arg2) // todo remind message

			return []string{}
		},
	},
}

var Bot *bot.Bot

func init() {
	var err error
	Bot, err = bot.NewBot(metadata, setting, pluginRules, commandRules)
	if err != nil {
		panic(err)
	}
}
