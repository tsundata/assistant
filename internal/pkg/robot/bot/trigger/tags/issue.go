package tags

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/vendors/github"
)

type Issue struct{}

func NewIssue() *Issue {
	return &Issue{}
}

func (t *Issue) Handle(ctx context.Context, comp component.Component, text string) {
	// get access token
	app, err := comp.Middle().GetAvailableApp(ctx, &pb.TextRequest{Text: github.ID})
	if err != nil {
		comp.GetLogger().Error(err)
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
		comp.GetLogger().Error(err)
		return
	}
	if *issue.ID == 0 {
		return
	}

	// send message
	err = comp.GetBus().Publish(ctx, enum.Message, event.MessageSendSubject, pb.Message{Text: fmt.Sprintf("Created Issue #%d %s", *issue.Number, *issue.HTMLURL)})
	if err != nil {
		comp.GetLogger().Error(err)
		return
	}
}
