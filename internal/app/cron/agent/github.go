package agent

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/vendors/github"
)

func FetchGithubStarred(b *rulebot.RuleBot) []string {
	// get access token
	app, err := b.MidClient.GetAvailableApp(context.Background(), &pb.TextRequest{Text: github.ID})
	if err != nil {
		return []string{}
	}
	accessToken := app.GetToken()
	if accessToken == "" {
		return []string{}
	}

	// data
	client := github.NewGithub("", "", "", accessToken)
	user, err := client.GetUser()
	if err != nil {
		return []string{}
	}
	if *user.Login != "" {
		repos, err := client.GetStarred(*user.Login)
		if err != nil {
			return []string{}
		}
		var result []string
		for _, item := range *repos {
			result = append(result, fmt.Sprintf("%s (%s)", *item.FullName, *item.HTMLURL))
		}
		return result
	}

	return []string{}
}
