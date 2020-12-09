package plugins

import (
	"context"
	"github.com/tsundata/assistant/internal/app/message/bot"
	"github.com/tsundata/assistant/internal/app/message/plugins/rules/cron"
	"github.com/tsundata/assistant/internal/app/message/plugins/rules/regex"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var regexRules = []regex.Rule{
	{
		Regex:       `qr (.*)`,
		HelpMessage: `Generate QR code`,
		ParseMessage: func(b *bot.Bot, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}

			txt := args[1]
			var reply string
			err := b.WebClient.Call(context.Background(), "Qr", &txt, &reply)
			if err != nil {
				return []string{"error call"}
			}

			return []string{
				reply,
			}
		},
	},
	{
		Regex:       `ut (\d+)`,
		HelpMessage: `Unix Timestamp`,
		ParseMessage: func(b *bot.Bot, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}

			utArg := args[1]
			tt, err := strconv.ParseInt(utArg, 10, 64)
			if err != nil {
				return []string{"error"}
			}

			t := time.Unix(tt, 0)

			return []string{
				t.String(),
			}
		},
	},
	{
		Regex:       `rand (\d+) (\d+)`,
		HelpMessage: `Unix Timestamp`,
		ParseMessage: func(b *bot.Bot, s string, args []string) []string {
			if len(args) != 3 {
				return []string{"error args"}
			}

			minArg := args[1]
			maxArg := args[2]
			min, err := strconv.Atoi(minArg)
			if err != nil {
				return []string{"error"}
			}
			max, err := strconv.Atoi(maxArg)
			if err != nil {
				return []string{"error"}
			}

			rand.Seed(time.Now().Unix())
			t := rand.Intn(max-min) + min

			return []string{
				strconv.Itoa(t),
			}
		},
	},
	{
		Regex:       `pwd (\d+)`,
		HelpMessage: `Generate Password`,
		ParseMessage: func(b *bot.Bot, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}

			lenArg := args[1]
			length, err := strconv.Atoi(lenArg)
			if err != nil {
				return []string{"error"}
			}

			pwd := utils.GeneratePassword(length, "lowercase|uppercase|numbers")

			return []string{
				pwd,
			}
		},
	},
	{
		Regex:       `subs list`,
		HelpMessage: `List subscribe`,
		ParseMessage: func(b *bot.Bot, s string, args []string) []string {
			var reply []string
			err := b.SubClient.Call(context.Background(), "List", nil, &reply)
			if err != nil {
				return []string{"error call"}
			}

			return reply
		},
	},
	{
		Regex:       `subs open (.*)`,
		HelpMessage: `Open subscribe`,
		ParseMessage: func(b *bot.Bot, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}
			var reply bool
			err := b.SubClient.Call(context.Background(), "Open", &args[1], &reply)
			if err != nil {
				return []string{"error call"}
			}
			if reply {
				return []string{"success"}
			}

			return []string{"failed"}
		},
	},
	{
		Regex:       `subs close (.*)`,
		HelpMessage: `Close subscribe`,
		ParseMessage: func(b *bot.Bot, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}
			var reply bool
			err := b.SubClient.Call(context.Background(), "Close", &args[1], &reply)
			if err != nil {
				return []string{"error call"}
			}
			if reply {
				return []string{"success"}
			}

			return []string{"failed"}
		},
	},
}

var cronRules = map[string]cron.Rule{
	"heartbeat": {
		"0 0 * * *",
		func() []model.Event {
			log.Println("cron " + time.Now().String())
			return []model.Event{
				{
					Data: model.EventData{Message: model.Message{
						Text: "Plugin Cron Heartbeat: " + time.Now().String(),
					}},
				},
			}
		},
	},
}

var Options = []bot.Option{
	bot.RegisterRuleset(regex.New(regexRules)),
	bot.RegisterRuleset(cron.New(cronRules)),
}
