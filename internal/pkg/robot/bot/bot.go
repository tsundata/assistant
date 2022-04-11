package bot

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

type Bot struct {
	Metadata
	SettingRule  []FieldItem
	WorkflowRule WorkflowRule
	CommandRule  []command.Rule
	ActionRule   []ActionRule
	FormRule     []FormRule
	TagRule      []TagRule
	plugin       []Plugin

	config *Config
	ctrl   *Controller
}

type Metadata struct {
	Name       string
	Identifier string
	Detail     string
	Avatar     string
}

type FieldItemType string

const (
	FieldItemTypeString FieldItemType = "string"
	FieldItemTypeInt    FieldItemType = "int"
	FieldItemTypeFloat  FieldItemType = "float"
	FieldItemTypeBool   FieldItemType = "bool"
)

type FieldItem struct {
	Key      string        `json:"key"`
	Type     FieldItemType `json:"type"`
	Required bool          `json:"required"`
	Value    interface{}   `json:"value"`
}

type WorkflowRule struct {
	Plugin  []PluginRule
	RunFunc ActFunc
}

type PluginRule struct {
	Name  string
	Param []interface{}
}

type ActionRule struct {
	ID         string
	Title      string
	OptionFunc map[string]ActFunc
}

type FormRule struct {
	ID         string
	Title      string
	Field      []FieldItem
	SubmitFunc ActFunc
}

type TagRule struct {
	Tag         string
	TriggerFunc ActFunc
}

type ActFunc func(context.Context, Context, component.Component) []pb.MsgPayload

func NewBot(metadata Metadata, settings []FieldItem,
	workflowRule WorkflowRule,
	commandsRule []command.Rule,
	actionRule []ActionRule,
	formRule []FormRule,
	tagRule []TagRule) (*Bot, error) {
	cfg := &Config{}
	b := &Bot{
		Metadata:     metadata,
		SettingRule:  settings,
		WorkflowRule: workflowRule,
		CommandRule:  commandsRule,
		ActionRule:   actionRule,
		FormRule:     formRule,
		TagRule:      tagRule,
		config:       cfg,
	}
	ctrl := &Controller{
		Instance:    b,
		Config:      cfg,
		PluginParam: make(map[string][]interface{}),
	}
	b.ctrl = ctrl

	// setup plugins
	err := SetupPlugins(ctrl, workflowRule.Plugin)
	if err != nil {
		return nil, err
	}
	b.plugin = cfg.Plugin

	// plugin chain
	var stack PluginHandler
	for i := len(b.plugin) - 1; i >= 0; i-- {
		stack = b.plugin[i](stack)
		b.config.RegisterHandler(stack)
	}
	b.config.PluginChain = stack

	return b, nil
}

func (b *Bot) RunPlugin(ctx context.Context, comp component.Component, input interface{}) (interface{}, error) {
	b.ctrl.Comp = comp
	if b.config.PluginChain != nil {
		return b.config.PluginChain.Run(ctx, b.ctrl, input)
	}
	return input, nil
}

func (b *Bot) Info() string {
	return fmt.Sprintf("bot:%s, %s", b.Name, b.Detail)
}
