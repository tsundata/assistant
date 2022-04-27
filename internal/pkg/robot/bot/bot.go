package bot

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

type Bot struct {
	Metadata
	SettingRule  []FieldItem
	EventHandler func(component.Component) error
	WorkflowRule WorkflowRule
	CommandRule  []command.Rule
	ActionRule   []ActionRule
	FormRule     []FormRule
	TagRule      []TagRule

	ctrl *Controller
}

func NewBot(metadata Metadata, settings []FieldItem, event func(component.Component) error,
	workflowRule WorkflowRule,
	commandsRule []command.Rule,
	actionRule []ActionRule,
	formRule []FormRule,
	tagRule []TagRule) (*Bot, error) {
	b := &Bot{
		Metadata:     metadata,
		SettingRule:  settings,
		EventHandler: event,
		WorkflowRule: workflowRule,
		CommandRule:  commandsRule,
		ActionRule:   actionRule,
		FormRule:     formRule,
		TagRule:      tagRule,
	}
	b.ctrl = &Controller{
		Instance: b,
	}

	return b, nil
}

func (b *Bot) RunPlugin(ctx context.Context, comp component.Component, input PluginValue) (PluginValue, error) {
	// setup
	plugin, params := SetupPlugins(b.WorkflowRule.Plugin)

	// run
	var err error
	b.ctrl.Comp = comp
	output := PluginValue{}
	for i, item := range plugin {
		output, err = item.Run(ctx, b.ctrl, params[i], input)
		fmt.Println("[robot] 		run plugin:", item.Name(), params[i], input, output, err)
		if err != nil {
			return PluginValue{}, err
		}
		input = output
	}

	return output, nil
}

func (b *Bot) Info() string {
	return fmt.Sprintf("bot:%s, %s", b.Name, b.Detail)
}
