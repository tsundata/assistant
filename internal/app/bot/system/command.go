package system

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/tsundata/assistant/api/pb"
	tags2 "github.com/tsundata/assistant/internal/pkg/robot/bot/trigger/tags"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/version"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var commandRules = []command.Rule{
	{
		Define: "version",
		Help:   `Version info`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			return []string{version.Info()}
		},
	},
	{
		Define: `qr [string]`,
		Help:   `Generate QR code`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}

			reply, err := comp.Middle().GetQrUrl(ctx, &pb.TextRequest{
				Text: tokens[1].Value.(string),
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
		Define: `ut [number]`,
		Help:   `Unix Timestamp`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			t := time.Unix(tokens[1].Value.(int64), 0)

			return []string{
				t.String(),
			}
		},
	},
	{
		Define: `rand [number] [number]`,
		Help:   `Unix Timestamp`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			min := tokens[1].Value.(int64)
			max := tokens[2].Value.(int64)

			nBing, err := rand.Int(rand.Reader, big.NewInt(max+1-min))
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			t := nBing.Int64() + min

			return []string{
				strconv.FormatInt(t, 10),
			}
		},
	},
	{
		Define: `pwd [number]`,
		Help:   `Generate Password`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			length := tokens[1].Value.(int64)

			pwd := util.RandString(int(length), "lowercase|uppercase|numbers")

			return []string{
				pwd,
			}
		},
	},
	{
		Define: "subs list",
		Help:   `List subscribe`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Middle().ListSubscribe(ctx, &pb.SubscribeRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			tableString := &strings.Builder{}
			if len(reply.Subscribe) > 0 {
				table := tablewriter.NewWriter(tableString)
				table.SetBorder(false)
				table.SetHeader([]string{"Name", "Subscribe"})
				for _, v := range reply.Subscribe {
					table.Append([]string{v.Name, util.BoolToString(v.State)})
				}
				table.Render()
			}
			if tableString.String() == "" {
				return []string{"empty subscript"}
			}

			return []string{tableString.String()}
		},
	},
	{
		Define: "subs open [string]",
		Help:   `Open subscribe`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}

			reply, err := comp.Middle().OpenSubscribe(ctx, &pb.SubscribeRequest{
				Text: tokens[2].Value.(string),
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `subs close [string]`,
		Help:   `Close subscribe`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}

			reply, err := comp.Middle().CloseSubscribe(ctx, &pb.SubscribeRequest{
				Text: tokens[2].Value.(string),
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `view [number]`, // todo sequence
		Help:   `View message`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Message() == nil {
				return []string{"empty client"}
			}

			id := tokens[1].Value.(int64)
			messageReply, err := comp.Message().Get(ctx, &pb.MessageRequest{Message: &pb.Message{Id: id}})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			if messageReply.Message.Sequence == 0 {
				return []string{"no message"}
			}

			return []string{messageReply.Message.GetText()}
		},
	},
	{
		Define: `run [number]`,
		Help:   `Run message`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Message() == nil {
				return []string{"empty client"}
			}

			id := tokens[1].Value.(int64)

			reply, err := comp.Message().Run(ctx, &pb.MessageRequest{Message: &pb.Message{Id: id}})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{reply.GetText()}
		},
	},
	{
		Define: `doc`,
		Help:   `Show action docs`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Chatbot() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Chatbot().ActionDoc(ctx, &pb.WorkflowRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			res := strings.Builder{}
			res.WriteString("Action:\n")
			res.WriteString(reply.GetText())
			res.WriteString("\n\nTag:\n")
			for k := range tags2.Tags() {
				res.WriteString("#")
				res.WriteString(k)
				res.WriteString("\n")
			}
			return []string{res.String()}
		},
	},
	{
		Define: `test`,
		Help:   `Test`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			_, err := comp.Message().Send(ctx, &pb.MessageRequest{Message: &pb.Message{Text: "test"}})
			if err != nil {
				return []string{err.Error()}
			}
			_, err = comp.Middle().GetSubscribeStatus(ctx, &pb.SubscribeRequest{Text: "example"})
			if err != nil {
				return []string{err.Error()}
			}
			_, err = comp.NLP().Pinyin(ctx, &pb.TextRequest{Text: "测试"})
			if err != nil {
				return []string{err.Error()}
			}
			_, err = comp.User().GetUser(ctx, &pb.UserRequest{User: &pb.User{Id: 1}})
			if err != nil {
				return []string{err.Error()}
			}
			_, err = comp.Chatbot().SyntaxCheck(ctx, &pb.WorkflowRequest{Text: "echo 1"})
			if err != nil {
				return []string{err.Error()}
			}
			return []string{"test done"}
		},
	},
	{
		Define: `stats`,
		Help:   `Stats Info`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Middle().GetStats(ctx, &pb.TextRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			return []string{reply.GetText()}
		},
	},
	{
		Define: `pinyin [string]`,
		Help:   "chinese pinyin conversion",
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.NLP() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.NLP().Pinyin(ctx, &pb.TextRequest{Text: tokens[1].Value.(string)})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if len(reply.GetText()) <= 0 {
				return []string{"failed"}
			}
			return []string{strings.Join(reply.GetText(), ", ")}
		},
	},
	{
		Define: `del [number]`,
		Help:   `Delete message`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			id := tokens[1].Value.(int64)

			_, err := comp.Message().Delete(ctx, &pb.MessageRequest{Message: &pb.Message{Id: id}})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{fmt.Sprintf("Deleted %d", id)}
		},
	},
	{
		Define: "cron list",
		Help:   `List cron`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Middle().ListCron(ctx, &pb.CronRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			tableString := &strings.Builder{}
			if len(reply.Cron) > 0 {
				table := tablewriter.NewWriter(tableString)
				table.SetBorder(false)
				table.SetHeader([]string{"Name", "IsCron"})
				for _, v := range reply.Cron {
					table.Append([]string{v.Name, util.BoolToString(v.State)})
				}
				table.Render()
			}
			if tableString.String() == "" {
				return []string{"empty cron"}
			}

			return []string{tableString.String()}
		},
	},
	{
		Define: "cron start [string]",
		Help:   `Start cron`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}

			reply, err := comp.Middle().StartCron(ctx, &pb.CronRequest{
				Text: tokens[2].Value.(string),
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `cron stop [string]`,
		Help:   `Stop cron`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}

			reply, err := comp.Middle().StopCron(ctx, &pb.CronRequest{
				Text: tokens[2].Value.(string),
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `webhook list`,
		Help:   `List webhook`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Chatbot() == nil {
				return []string{"empty client"}
			}

			reply, err := comp.Chatbot().ListWebhook(ctx, &pb.WorkflowRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			var result strings.Builder
			for _, flag := range reply.Flag {
				result.WriteString("/webhook/")
				result.WriteString(flag)
				result.WriteString("\n")
			}
			if result.Len() > 0 {
				return []string{result.String()}
			}

			return []string{"failed"}
		},
	},
}
