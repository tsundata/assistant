package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/vendors/dropbox"
	"time"
)

func Backup(ctx rulebot.IContext) []result.Result {
	if ctx.Middle() == nil || ctx.Message() == nil || ctx.Todo() == nil {
		return []result.Result{result.EmptyResult()}
	}
	ctxB := context.Background()
	app, err := ctx.Middle().GetAvailableApp(ctxB, &pb.TextRequest{Text: dropbox.ID})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}
	accessToken := app.GetToken()
	if accessToken == "" {
		return []result.Result{result.ErrorResult(errors.New("backup: dropbox access token is empty"))}
	}

	// messages
	messagesReply, err := ctx.Message().List(ctxB, &pb.MessageRequest{})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}

	// apps
	appsReply, err := ctx.Middle().GetApps(ctxB, &pb.TextRequest{})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}

	// credentials
	credentialsReply, err := ctx.Middle().GetCredentials(ctxB, &pb.TextRequest{})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}

	// todos
	todosReply, err := ctx.Todo().GetTodos(ctxB, &pb.TodoRequest{})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}

	data := map[string]interface{}{
		"message":     messagesReply.Messages,
		"apps":        appsReply.Apps,
		"credentials": credentialsReply.Credentials,
		"todos":       todosReply.Todos,
	}
	d, err := json.Marshal(data)
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}

	// upload
	client := dropbox.NewDropbox("", "", "", accessToken)
	path := fmt.Sprintf("/backup/assistant_%s.json", time.Now().Format("2006-01-02_15:04:05"))
	err = client.Upload(path, bytes.NewReader(d))
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}

	return []result.Result{result.MessageResult(fmt.Sprintf("backed up in %s", time.Now()))}
}
