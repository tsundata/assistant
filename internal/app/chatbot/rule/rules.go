package rule

import (
	"context"
	"errors"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/version"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var rules = []Rule{
	{
		Regex:       `version`,
		HelpMessage: `Version info`,
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			return []string{version.Info()}
		},
	},
	{
		Regex:       `menu`,
		HelpMessage: `Show menu`,
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if ctx.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := ctx.Middle().GetMenu(context.Background(), &pb.TextRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			if reply.GetText() == "" {
				return []string{"empty menu"}
			}

			return []string{reply.GetText()}
		},
	},
	{
		Regex:       `qr\s+(.*)`,
		HelpMessage: `Generate QR code`,
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if ctx.Middle() == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}

			txt := args[1]
			reply, err := ctx.Middle().GetQrUrl(context.Background(), &pb.TextRequest{
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
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
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
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
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
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}

			lenArg := args[1]
			length, err := strconv.Atoi(lenArg)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			pwd := util.GeneratePassword(length, "lowercase|uppercase|numbers")

			return []string{
				pwd,
			}
		},
	},
	{
		Regex:       `subs\s+list`,
		HelpMessage: `List subscribe`,
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if ctx.Subscribe() == nil {
				return []string{"empty client"}
			}
			reply, err := ctx.Subscribe().List(context.Background(), &pb.SubscribeRequest{})
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
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if ctx.Subscribe() == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}

			reply, err := ctx.Subscribe().Open(context.Background(), &pb.SubscribeRequest{
				Text: args[1],
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
		Regex:       `subs\s+close\s+(.*)`,
		HelpMessage: `Close subscribe`,
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if ctx.Subscribe() == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}

			reply, err := ctx.Subscribe().Close(context.Background(), &pb.SubscribeRequest{
				Text: args[1],
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
		Regex:       `view\s+(\d+)`,
		HelpMessage: `View message`,
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if ctx.Message() == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}

			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return []string{"error args"}
			}
			messageReply, err := ctx.Message().Get(context.Background(), &pb.MessageRequest{Id: id})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			if messageReply.Id == 0 {
				return []string{"no message"}
			}

			return []string{messageReply.GetText()}
		},
	},
	{
		Regex:       `run\s+(\d+)`,
		HelpMessage: `Run message`,
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if ctx.Message() == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}

			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return []string{"error args"}
			}

			reply, err := ctx.Message().Run(context.Background(), &pb.MessageRequest{Id: id})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{reply.GetText()}
		},
	},
	{
		Regex:       `doc`,
		HelpMessage: `Show action docs`,
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if ctx.Workflow() == nil {
				return []string{"empty client"}
			}
			reply, err := ctx.Workflow().ActionDoc(context.Background(), &pb.WorkflowRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			return []string{reply.GetText()}
		},
	},
	{
		Regex:       `test`,
		HelpMessage: `Test`,
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if ctx.Storage() == nil {
				return []string{"empty client"}
			}
			// upload
			f, err := os.Open("./README.md")
			if err != nil {
				return []string{"error: " + err.Error()}
			}

			buf := make([]byte, 1024)
			uc, err := ctx.Storage().UploadFile(context.Background())
			if err != nil {
				return []string{"error: " + err.Error()}
			}

			err = uc.Send(&pb.FileRequest{Data: &pb.FileRequest_Info{Info: &pb.FileInfo{FileType: "md"}}})
			if err != nil {
				return []string{"error: " + err.Error()}
			}

			for {
				n, err := f.Read(buf)
				if errors.Is(err, io.EOF) {
					break
				}
				if err != nil {
					return []string{"error: " + err.Error()}
				}
				err = uc.Send(&pb.FileRequest{Data: &pb.FileRequest_Chuck{Chuck: buf[:n]}})
				if err != nil {
					return []string{"error: " + err.Error()}
				}
			}

			_, err = uc.CloseAndRecv()
			if err != nil {
				return []string{"error: " + err.Error()}
			}

			return []string{"test done"}
		},
	},
	{
		Regex:       `stats`,
		HelpMessage: `Stats Info`,
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if ctx.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := ctx.Middle().GetStats(context.Background(), &pb.TextRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			return []string{reply.GetText()}
		},
	},
	{
		Regex:       `todo\s+(.*)`,
		HelpMessage: "Todo something",
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if ctx.Todo() == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}
			reply, err := ctx.Todo().CreateTodo(context.Background(), &pb.TodoRequest{
				Content: args[1],
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if !reply.GetState() {
				return []string{"failed"}
			}
			return []string{"success"}
		},
	},
	{
		Regex:       `role`,
		HelpMessage: "Role info",
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if ctx.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := ctx.Middle().GetRoleImageUrl(context.Background(), &pb.TextRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			return []string{reply.GetText()}
		},
	},
	{
		Regex:       `pinyin\s+(.*)`,
		HelpMessage: "chinese pinyin conversion",
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if ctx.NLP() == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}
			reply, err := ctx.NLP().Pinyin(context.Background(), &pb.TextRequest{Text: args[1]})
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
		Regex:       `remind\s+(\w+)\s+(\w+)`,
		HelpMessage: `Remind something`,
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if len(args) != 3 {
				return []string{"error args"}
			}

			arg1 := args[1]
			arg2 := args[2]
			fmt.Println(arg1, arg2) // todo remind message

			return []string{}
		},
	},
	{
		Regex:       `del\s+(\d+)`,
		HelpMessage: `Delete message`,
		ParseMessage: func(ctx rulebot.IContext, s string, args []string) []string {
			if len(args) != 2 {
				return []string{"error args"}
			}

			idStr := args[1]
			id, err := strconv.Atoi(idStr)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			fmt.Println(id) // todo delete message

			return []string{}
		},
	},
}

var Options = []rulebot.Option{
	rulebot.RegisterRuleset(New(rules)),
}
