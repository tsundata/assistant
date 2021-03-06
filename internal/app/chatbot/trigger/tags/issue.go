package tags

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/vendors/github"
)

type Issue struct{}

func NewIssue() *Issue {
	return &Issue{}
}

func (t *Issue) Handle(ctx *ctx.Context, text string) {
	// get access token
	app, err := ctx.Middle.GetAvailableApp(context.Background(), &pb.TextRequest{Text: github.ID})
	if err != nil {
		ctx.Logger.Error(err)
		return
	}
	accessToken := app.GetToken()
	if accessToken == "" {
		return
	}

	// get user
	client := github.NewGithub("", "", "", accessToken)
	user, err := client.GetUser()
	if err != nil {
		return
	}
	if *user.Login == "" {
		return
	}

	// create issue
	issue, err := client.CreateIssue(*user.Login, "assistant", github.Issue{Title: &text})
	if err != nil {
		ctx.Logger.Error(err)
		return
	}
	if *issue.ID == 0 {
		return
	}

	// send message
	err = ctx.Bus.Publish(event.SendMessageSubject, model.Message{Text: fmt.Sprintf("Created Issue #%d %s", *issue.Number, *issue.HTMLURL)})
	if err != nil {
		ctx.Logger.Error(err)
		return
	}
}
