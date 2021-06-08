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
	"github.com/tsundata/assistant/internal/pkg/transports/rpc/rpcclient"
	"github.com/tsundata/assistant/internal/pkg/vendors/dropbox"
	"time"
)

func Backup(b *rulebot.RuleBot) []result.Result {
	ctx := context.Background()
	app, err := rpcclient.GetMiddleClient(b.Client).GetAvailableApp(ctx, &pb.TextRequest{Text: dropbox.ID})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}
	accessToken := app.GetToken()
	if accessToken == "" {
		return []result.Result{result.ErrorResult(errors.New("backup: dropbox access token is empty"))}
	}

	// messages
	messagesReply, err := rpcclient.GetMessageClient(b.Client).List(ctx, &pb.MessageRequest{})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}

	// apps
	appsReply, err := rpcclient.GetMiddleClient(b.Client).GetApps(ctx, &pb.TextRequest{})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}

	// credentials
	credentialsReply, err := rpcclient.GetMiddleClient(b.Client).GetCredentials(ctx, &pb.TextRequest{})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}

	data := map[string]interface{}{
		"message":     messagesReply.Messages,
		"apps":        appsReply.Apps,
		"credentials": credentialsReply.Credentials,
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
