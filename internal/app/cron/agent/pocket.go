package agent

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/rpcclient"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors/pocket"
)

func FetchPocket(b *rulebot.RuleBot) []result.Result {
	// get consumer key
	reply, err := rpcclient.GetMiddleClient(b.Client).GetCredential(context.Background(), &pb.CredentialRequest{Type: pocket.ID})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}
	consumerKey := ""
	for _, item := range reply.GetContent() {
		if item.Key == pocket.ClientIdKey {
			consumerKey = item.Value
		}
	}
	if consumerKey == "" {
		return []result.Result{result.EmptyResult()}
	}

	// get access token
	app, err := rpcclient.GetMiddleClient(b.Client).GetAvailableApp(context.Background(), &pb.TextRequest{Text: pocket.ID})
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}
	accessToken := app.GetToken()
	if accessToken == "" {
		return []result.Result{result.EmptyResult()}
	}

	// data
	client := pocket.NewPocket(consumerKey, "", "", accessToken)
	resp, err := client.Retrieve(10)
	if err != nil {
		return []result.Result{result.ErrorResult(err)}
	}
	if resp.Status == 0 {
		return []result.Result{result.EmptyResult()}
	}

	var r []result.Result
	for _, item := range resp.List {
		r = append(r, result.Result{
			ID:   util.SHA1(item.ResolvedUrl),
			Kind: result.Url,
			Content: map[string]string{
				"title": item.ResolvedTitle,
				"url":   item.ResolvedUrl,
			},
		})
	}
	return r
}
