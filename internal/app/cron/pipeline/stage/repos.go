package stage

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/vendors/github"
)

func Repos(ctx rulebot.IContext, in result.Result) result.Result {
	if in.Kind == result.Repos {
		if data, ok := in.Content.(map[string]string); ok {
			// get access token
			app, err := ctx.Middle().GetAvailableApp(context.Background(), &pb.TextRequest{Text: github.ID})
			if err != nil {
				return result.ErrorResult(err)
			}
			accessToken := app.GetToken()
			if accessToken == "" {
				return result.EmptyResult()
			}

			client := github.NewGithub("", "", "", accessToken)
			repo, err := client.GetRepository(data["owner"], data["repo"])
			if err != nil {
				return result.ErrorResult(err)
			}

			return result.MessageResult(fmt.Sprintf("Repo: %s \nStar: %d\nFork: %d\nWatch: %d", *repo.FullName, *repo.StargazersCount, *repo.ForksCount, *repo.WatchersCount))
		}
	}
	return result.EmptyResult()
}
