package tags

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/vendors/github"
)

type Project struct{}

func NewProject() *Project {
	return &Project{}
}

func (t *Project) Handle(ctx context.Context, comp *ctx.Component, text string) {
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
		comp.Logger.Error(err)
		return
	}
	if *user.Login == "" {
		return
	}

	// get projects
	projects, err := client.GetUserProjects(*user.Login)
	if err != nil {
		comp.Logger.Error(err)
		return
	}
	if len(*projects) == 0 {
		return
	}

	// get columns
	columns, err := client.GetProjectColumns(*(*projects)[0].ID)
	if err != nil {
		comp.Logger.Error(err)
		return
	}
	if len(*columns) == 0 {
		return
	}

	// create card
	card, err := client.CreateCard(*(*columns)[0].ID, github.ProjectCard{Note: &text})
	if err != nil {
		comp.Logger.Error(err)
		return
	}
	if *card.ID == 0 {
		return
	}

	// send message
	err = comp.Bus.Publish(ctx, event.MessageSendSubject, pb.Message{Text: fmt.Sprintf("Created Project Card #%d", *card.ID)})
	if err != nil {
		comp.Logger.Error(err)
		return
	}
}
