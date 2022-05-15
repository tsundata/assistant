package agent

import (
	"context"
	"crypto/sha1" // #nosec
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors/pocket"
)

func FetchPocket(ctx context.Context, comp component.Component) []result.Result {
	if comp.Middle() == nil {
		return []result.Result{result.EmptyResult()}
	}
	// get consumer key
	reply, err := comp.Middle().GetCredential(ctx, &pb.CredentialRequest{Type: pocket.ID})
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
	app, err := comp.Middle().GetAvailableApp(ctx, &pb.TextRequest{Text: pocket.ID})
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
	for i := range resp.List {
		item := resp.List[i]
		s := sha1.New()                              // #nosec
		s.Write(util.StringToByte(item.ResolvedUrl)) // #nosec
		bs := s.Sum(nil)
		r = append(r, result.Result{
			ID:   util.ByteToString(bs),
			Kind: result.Url,
			Content: map[string]string{
				"title": item.ResolvedTitle,
				"url":   item.ResolvedUrl,
			},
		})
	}
	return r
}
