package plugins

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
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
			reply, err := (*b.MidClient).Qr(context.Background(), &pb.Text{
				Text: txt,
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{
				reply.GetText(),
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
				return []string{"error call: " + err.Error()}
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
				return []string{"error call: " + err.Error()}
			}
			max, err := strconv.Atoi(maxArg)
			if err != nil {
				return []string{"error call: " + err.Error()}
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
				return []string{"error call: " + err.Error()}
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
			reply, err := (*b.SubClient).List(context.Background(), &pb.SubscribeRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return reply.GetText()
		},
	},
	{
		Regex:       `subs open (.*)`,
		HelpMessage: `Open subscribe`,
		ParseMessage: func(b *bot.Bot, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}

			reply, err := (*b.SubClient).Open(context.Background(), &pb.SubscribeRequest{
				Text: args[1],
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.State {
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

			reply, err := (*b.SubClient).Open(context.Background(), &pb.SubscribeRequest{
				Text: args[1],
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.State {
				return []string{"success"}
			}

			return []string{"failed"}
		},
	},
}

var cronRules = map[string]cron.Rule{
	"heartbeat": {
		When: "0 0 * * *",
		Action: func() []model.Event {
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
