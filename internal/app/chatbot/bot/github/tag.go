package github

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/vendors/github"
)

var tagRules = []bot.TagRule{
	{
		Tag: "issue",
		TriggerFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
			text := botCtx.Message.GetText()
			// get access token
			app, err := comp.Middle().GetAvailableApp(ctx, &pb.TextRequest{Text: github.ID})
			if err != nil {
				comp.GetLogger().Error(err)
				return nil
			}
			accessToken := app.GetToken()
			if accessToken == "" {
				return nil
			}

			// get user
			client := github.NewGithub("", "", "", accessToken)
			user, err := client.GetUser()
			if err != nil {
				return nil
			}
			if *user.Login == "" {
				return nil
			}

			// create issue
			issue, err := client.CreateIssue(*user.Login, "assistant", github.Issue{Title: &text})
			if err != nil {
				comp.GetLogger().Error(err)
				return nil
			}
			if *issue.ID == 0 {
				return nil
			}

			// send message
			err = comp.GetBus().Publish(ctx, enum.Message, event.MessageSendSubject, pb.Message{
				Text: fmt.Sprintf("Created Issue #%d %s", *issue.Number, *issue.HTMLURL)}) //todo
			if err != nil {
				comp.GetLogger().Error(err)
				return nil
			}
			return nil
		},
	},
	{
		Tag: "project",
		TriggerFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
			text := botCtx.Message.GetText()
			// get access token
			app, err := comp.Middle().GetAvailableApp(ctx, &pb.TextRequest{Text: github.ID})
			if err != nil {
				comp.GetLogger().Error(err)
				return nil
			}
			accessToken := app.GetToken()
			if accessToken == "" {
				return nil
			}

			// get user
			client := github.NewGithub("", "", "", accessToken)
			user, err := client.GetUser()
			if err != nil {
				comp.GetLogger().Error(err)
				return nil
			}
			if *user.Login == "" {
				return nil
			}

			// get projects
			projects, err := client.GetUserProjects(*user.Login)
			if err != nil {
				comp.GetLogger().Error(err)
				return nil
			}
			if len(*projects) == 0 {
				return nil
			}

			// get columns
			columns, err := client.GetProjectColumns(*(*projects)[0].ID)
			if err != nil {
				comp.GetLogger().Error(err)
				return nil
			}
			if len(*columns) == 0 {
				return nil
			}

			// create card
			card, err := client.CreateCard(*(*columns)[0].ID, github.ProjectCard{Note: &text})
			if err != nil {
				comp.GetLogger().Error(err)
				return nil
			}
			if *card.ID == 0 {
				return nil
			}

			// send message
			err = comp.GetBus().Publish(ctx, enum.Message, event.MessageSendSubject, pb.Message{
				Text: fmt.Sprintf("Created Project Card #%d", *card.ID)}) //todo
			if err != nil {
				comp.GetLogger().Error(err)
				return nil
			}
			return nil
		},
	},
}
