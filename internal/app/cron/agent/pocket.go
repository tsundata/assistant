package agent

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/vendors/pocket"
)

func FetchPocket(b *rulebot.RuleBot) []string {
	// get consumer key
	reply, err := b.MidClient.GetCredential(context.Background(), &pb.CredentialRequest{Type: pocket.ID})
	if err != nil {
		return []string{}
	}
	consumerKey := ""
	for _, item := range reply.GetContent() {
		if item.Key == pocket.ClientIdKey {
			consumerKey = item.Value
		}
	}
	if consumerKey == "" {
		return []string{}
	}

	// get access token
	app, err := b.MidClient.GetAvailableApp(context.Background(), &pb.TextRequest{Text: pocket.ID})
	if err != nil {
		return []string{}
	}
	accessToken := app.GetToken()
	if accessToken == "" {
		return []string{}
	}

	// data
	client := pocket.NewPocket(consumerKey, "", "", accessToken)
	resp, err := client.Retrieve(10)
	if err != nil {
		return []string{}
	}

	if resp.Status > 0 {
		var result []string
		for _, item := range resp.List {
			result = append(result, fmt.Sprintf("%s (%s)", item.ResolvedTitle, item.ResolvedUrl))
		}
		return result
	}

	return []string{}
}
