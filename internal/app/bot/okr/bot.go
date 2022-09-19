package okr

import (
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

const (
	CreateObjectiveFormID      = "create_objective"
	UpdateObjectiveFormID      = "update_objective"
	CreateKeyResultFormID      = "create_key_result"
	UpdateKeyResultFormID      = "Update_key_result"
	CreateKeyResultValueFormID = "create_key_result_value"
)

var metadata = bot.Metadata{
	Name:       "Okr",
	Identifier: enum.OkrBot,
	Detail:     "",
	Avatar:     "",
}

var workflowRules bot.WorkflowRule

var Bot *bot.Bot

func init() {
	var err error
	Bot, err = bot.NewBot(metadata, nil, eventHandler, workflowRules, commandRules, nil, formRules, nil)
	if err != nil {
		panic(err)
	}
}
