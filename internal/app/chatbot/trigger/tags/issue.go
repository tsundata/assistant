package tags

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/vendors/github"
)

type Issue struct{}

func NewIssue() *Issue {
	return &Issue{}
}

func (t *Issue) Handle(ctx context.Context, comp *ctx.Component, text string) {
	// get access token
	app, err := comp.Middle.GetAvailableApp(ctx, &pb.TextRequest{Text: github.ID})
	if err != nil {
		comp.Logger.Error(err)
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
		comp.Logger.Error(err)
		return
	}
	if *issue.ID == 0 {
		return
	}

	// send message
	err = comp.Bus.Publish(ctx, event.MessageSendSubject, pb.Message{Text: fmt.Sprintf("Created Issue #%d %s", *issue.Number, *issue.HTMLURL)})
	if err != nil {
		comp.Logger.Error(err)
		return
	}
}
