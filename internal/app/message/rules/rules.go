package rules

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/tsundata/assistant/internal/pkg/version"
	"math/rand"
	"strconv"
	"time"
)

var rules = []Rule{
	{
		Regex:       `version`,
		HelpMessage: `Version info`,
		ParseMessage: func(b *rulebot.RuleBot, s string, args []string) []string {
			return []string{version.Info()}
		},
	},
	{
		Regex:       `menu`,
		HelpMessage: `Show menu`,
		ParseMessage: func(b *rulebot.RuleBot, s string, args []string) []string {
			reply, err := b.MidClient.GetMenu(context.Background(), &pb.TextRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			if reply.GetText() == "" {
				return []string{"empty subscript"}
			}

			return []string{reply.GetText()}
		},
	},
	{
		Regex:       `qr\s+(.*)`,
		HelpMessage: `Generate QR code`,
		ParseMessage: func(b *rulebot.RuleBot, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}

			txt := args[1]
			reply, err := b.MidClient.GetQrUrl(context.Background(), &pb.TextRequest{
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
		Regex:       `ut\s+(\d+)`,
		HelpMessage: `Unix Timestamp`,
		ParseMessage: func(b *rulebot.RuleBot, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}

			tt, err := strconv.ParseInt(args[1], 10, 64)
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
		Regex:       `rand\s+(\d+)\s+(\d+)`,
		HelpMessage: `Unix Timestamp`,
		ParseMessage: func(b *rulebot.RuleBot, s string, args []string) []string {
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
		Regex:       `pwd\s+(\d+)`,
		HelpMessage: `Generate Password`,
		ParseMessage: func(b *rulebot.RuleBot, s string, args []string) []string {
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
		Regex:       `subs\s+list`,
		HelpMessage: `List subscribe`,
		ParseMessage: func(b *rulebot.RuleBot, s string, args []string) []string {
			reply, err := b.SubClient.List(context.Background(), &pb.SubscribeRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			if reply.GetText() == nil {
				return []string{"empty subscript"}
			}

			return reply.GetText()
		},
	},
	{
		Regex:       `subs\s+open\s+(.*)`,
		HelpMessage: `Open subscribe`,
		ParseMessage: func(b *rulebot.RuleBot, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}

			reply, err := b.SubClient.Open(context.Background(), &pb.SubscribeRequest{
				Text: args[1],
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"success"}
			}

			return []string{"failed"}
		},
	},
	{
		Regex:       `subs\s+close\s+(.*)`,
		HelpMessage: `Close subscribe`,
		ParseMessage: func(b *rulebot.RuleBot, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}

			reply, err := b.SubClient.Close(context.Background(), &pb.SubscribeRequest{
				Text: args[1],
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"success"}
			}

			return []string{"failed"}
		},
	},
	{
		Regex:       `view\s+(\d+)`,
		HelpMessage: `View message`,
		ParseMessage: func(b *rulebot.RuleBot, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}

			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return []string{"error args"}
			}
			messageReply, err := b.MsgClient.Get(context.Background(), &pb.MessageRequest{Id: id})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{messageReply.GetText()}
		},
	},
	{
		Regex:       `run\s+(\d+)`,
		HelpMessage: `Run message`,
		ParseMessage: func(b *rulebot.RuleBot, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}

			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return []string{"error args"}
			}

			reply, err := b.MsgClient.Run(context.Background(), &pb.MessageRequest{Id: id})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{reply.GetText()}
		},
	},
	{
		Regex:       `doc`,
		HelpMessage: `Show action docs`,
		ParseMessage: func(b *rulebot.RuleBot, s string, args []string) []string {
			reply, err := b.WfClient.ActionDoc(context.Background(), &pb.WorkflowRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			return []string{reply.GetText()}
		},
	},
	{
		Regex:       `test`,
		HelpMessage: `Test`,
		ParseMessage: func(b *rulebot.RuleBot, s string, args []string) []string {
			return []string{"test done"}
		},
	},
	{
		Regex:       `stats`,
		HelpMessage: `Stats Info`,
		ParseMessage: func(b *rulebot.RuleBot, s string, args []string) []string {
			reply, err := b.MidClient.GetStats(context.Background(), &pb.TextRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			return []string{reply.GetText()}
		},
	},
}

var Options = []rulebot.Option{
	rulebot.RegisterRuleset(New(rules)),
}
