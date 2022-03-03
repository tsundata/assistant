package bot

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
)

type Bot struct {
	Metadata
	SettingRule []SettingField
	PluginRule  []PluginRule
	plugin      []Plugin

	config *Config
	ctrl   *Controller
}

type Metadata struct {
	Name       string
	Identifier string
	Detail     string
	Avatar     string
}

type SettingItemType string

const (
	SettingItemTypeString SettingItemType = "string"
	SettingItemTypeInt    SettingItemType = "int"
	SettingItemTypeFloat  SettingItemType = "float"
	SettingItemTypeBool   SettingItemType = "bool"
)

type SettingField struct {
	Key      string          `json:"key"`
	Type     SettingItemType `json:"type"`
	Required bool            `json:"required"`
	Value    interface{}     `json:"value"`
}

type PluginRule struct {
	Name  string
	Param []interface{}
}

func NewBot(metadata Metadata, settings []SettingField, rules []PluginRule) (*Bot, error) {
	cfg := &Config{}
	b := &Bot{
		Metadata:    metadata,
		SettingRule: settings,
		PluginRule:  rules,
		config:      cfg,
	}
	ctrl := &Controller{
		Instance:    b,
		Config:      cfg,
		PluginParam: make(map[string][]interface{}),
	}
	b.ctrl = ctrl

	// setup plugins
	err := SetupPlugins(ctrl, rules)
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

func (b *Bot) Run(ctx context.Context, input interface{}) (interface{}, error) {
	b.ctrl.Ctx = ctx
	if b.config.PluginChain != nil {
		return b.config.PluginChain.Run(b.ctrl, input)
	}
	return input, nil
}

func (b *Bot) Info() string {
	return fmt.Sprintf("bot:%s, %s", b.Name, b.Detail)
}

func RegisterBot(ctx context.Context, bus event.Bus, bots ...*Bot) error {
	for _, item := range bots {
		err := bus.Publish(ctx, enum.Chatbot, event.BotRegisterSubject, pb.Bot{
			Name:       item.Name,
			Identifier: item.Identifier,
			Detail:     item.Detail,
			Avatar:     item.Avatar,
			Extend:     "",
		})
		if err != nil {
			return err
		}
	}
	return nil
}
