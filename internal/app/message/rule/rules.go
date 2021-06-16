package rule

import (
	"context"
	"errors"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/rpcclient"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/version"
	"io"
	"log"
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
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
			return []string{version.Info()}
		},
	},
	{
		Regex:       `menu`,
		HelpMessage: `Show menu`,
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
			if b.Client == nil {
				return []string{"empty client"}
			}
			reply, err := rpcclient.GetMiddleClient(b.Client).GetMenu(context.Background(), &pb.TextRequest{})
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
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
			if b.Client == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}

			txt := args[1]
			reply, err := rpcclient.GetMiddleClient(b.Client).GetQrUrl(context.Background(), &pb.TextRequest{
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
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
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
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
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
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
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
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
			if b.Client == nil {
				return []string{"empty client"}
			}
			reply, err := rpcclient.GetSubscribeClient(b.Client).List(context.Background(), &pb.SubscribeRequest{})
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
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
			if b.Client == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}

			reply, err := rpcclient.GetSubscribeClient(b.Client).Open(context.Background(), &pb.SubscribeRequest{
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
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
			if b.Client == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}

			reply, err := rpcclient.GetSubscribeClient(b.Client).Close(context.Background(), &pb.SubscribeRequest{
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
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
			if b.Client == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}

			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return []string{"error args"}
			}
			messageReply, err := rpcclient.GetMessageClient(b.Client).Get(context.Background(), &pb.MessageRequest{Id: id})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{messageReply.GetText()}
		},
	},
	{
		Regex:       `run\s+(\d+)`,
		HelpMessage: `Run message`,
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
			if b.Client == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}

			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return []string{"error args"}
			}

			reply, err := rpcclient.GetMessageClient(b.Client).Run(context.Background(), &pb.MessageRequest{Id: id})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{reply.GetText()}
		},
	},
	{
		Regex:       `doc`,
		HelpMessage: `Show action docs`,
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
			if b.Client == nil {
				return []string{"empty client"}
			}
			reply, err := rpcclient.GetWorkflowClient(b.Client).ActionDoc(context.Background(), &pb.WorkflowRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			return []string{reply.GetText()}
		},
	},
	{
		Regex:       `test`,
		HelpMessage: `Test`,
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
			if b.Client == nil {
				return []string{"empty client"}
			}
			// upload
			f, err := os.Open("./README.md")
			if err != nil {
				return []string{"error: " + err.Error()}
			}

			buf := make([]byte, 1024)
			uc, err := rpcclient.GetStorageClient(b.Client).UploadFile(context.Background())
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

			state, err := uc.CloseAndRecv()
			if err != nil {
				return []string{"error: " + err.Error()}
			}
			log.Println(state.GetPath())

			return []string{"test done"}
		},
	},
	{
		Regex:       `stats`,
		HelpMessage: `Stats Info`,
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
			if b.Client == nil {
				return []string{"empty client"}
			}
			reply, err := rpcclient.GetMiddleClient(b.Client).GetStats(context.Background(), &pb.TextRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			return []string{reply.GetText()}
		},
	},
	{
		Regex:       `todo\s+(.*)`,
		HelpMessage: "Todo something",
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
			if b.Client == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}
			reply, err := rpcclient.GetTodoClient(b.Client).CreateTodo(context.Background(), &pb.TodoRequest{
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
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
			if b.Client == nil {
				return []string{"empty client"}
			}
			reply, err := rpcclient.GetUserClient(b.Client).GetRole(context.Background(), &pb.RoleRequest{Id: model.SuperUserID})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetRole().Level <= 0 {
				return []string{"failed"}
			}
			return []string{fmt.Sprintf("%+v", reply.GetRole())}
		},
	},
	{
		Regex:       `pinyin\s+(.*)`,
		HelpMessage: "chinese pinyin conversion",
		ParseMessage: func(b *rulebot.Context, s string, args []string) []string {
			if b.Client == nil {
				return []string{"empty client"}
			}
			if len(args) != 2 {
				return []string{"error args"}
			}
			reply, err := rpcclient.GetNLPClient(b.Client).Pinyin(context.Background(), &pb.TextRequest{Text: args[1]})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if len(reply.GetText()) <= 0 {
				return []string{"failed"}
			}
			return []string{strings.Join(reply.GetText(), ", ")}
		},
	},
}

var Options = []rulebot.Option{
	rulebot.RegisterRuleset(New(rules)),
}
