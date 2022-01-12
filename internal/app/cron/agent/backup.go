package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/vendors/dropbox"
	"time"
)

func Backup(ctx context.Context, comp rulebot.IComponent) []result.Result {
	if comp.Middle() == nil || comp.Message() == nil || comp.Todo() == nil {
		return []result.Result{result.EmptyResult()}
	}
	app, err := comp.Middle().GetAvailableApp(ctx, &pb.TextRequest{Text: dropbox.ID})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}
	accessToken := app.GetToken()
	if accessToken == "" {
		return []result.Result{result.ErrorResult(errors.New("backup: dropbox access token is empty"))}
	}

	// messages
	messagesReply, err := comp.Message().List(ctx, &pb.GetMessagesRequest{}) // todo fixme
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}

	// apps
	appsReply, err := comp.Middle().GetApps(ctx, &pb.TextRequest{})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}

	// credentials
	credentialsReply, err := comp.Middle().GetCredentials(ctx, &pb.TextRequest{})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}

	// todos
	todosReply, err := comp.Todo().GetTodos(ctx, &pb.TodoRequest{})
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
