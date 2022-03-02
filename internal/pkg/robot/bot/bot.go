package bot

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
)

type Bot struct {
	Metadata
	SettingRule []SettingField
	PluginRule  []PluginRule
	plugin      []Plugin
	config      *Config
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
		Instance: b,
		Config:   cfg,
	}

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

func (b *Bot) Run(ctx context.Context) error {
	_, err := b.config.PluginChain.Run(ctx, nil)
	return err
}

func (b *Bot) Info() string {
	return fmt.Sprintf("bot:%s, %s", b.Name, b.Detail)
}

func RegisterBot(ctx context.Context, bus event.Bus, bots ...*Bot) error {
	for _, item := range bots {
		err := bus.Publish(ctx, event.BotRegisterSubject, pb.Bot{
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
