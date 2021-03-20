package tags

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/vendors/github"
)

type Todo struct{}

func NewTodo() *Todo {
	return &Todo{}
}

func (t *Todo) Handle(ctx *ctx.Context, text string) {
	// get access token
	app, err := ctx.MidClient.GetAvailableApp(context.Background(), &pb.TextRequest{Text: github.ID})
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

	// get projects
	projects, err := client.GetUserProjects(*user.Login)
	if err != nil {
		ctx.Logger.Error(err)
		return
	}
	if len(*projects) == 0 {
		return
	}

	// get columns
	columns, err := client.GetProjectColumns(*(*projects)[0].ID)
	if err != nil {
		ctx.Logger.Error(err)
		return
	}
	if len(*columns) == 0 {
		return
	}

	// create card
	card, err := client.CreateCard(*(*columns)[0].ID, github.ProjectCard{Note: &text})
	if err != nil {
		ctx.Logger.Error(err)
		return
	}
	if *card.ID == 0 {
		return
	}

	// send message
	_, err = ctx.MsgClient.Send(context.Background(), &pb.MessageRequest{Text: fmt.Sprintf("Created Todo Card #%d", *card.ID)})
	if err != nil {
		ctx.Logger.Error(err)
		return
	}
}
