package bot

import (
	"context"
	"fmt"
)

type Bot struct {
	Metadata
	Setting []SettingItem
	plugin  []Plugin
	config  *Config
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

type SettingItem struct {
	Key      string          `json:"key"`
	Type     SettingItemType `json:"type"`
	Required bool            `json:"required"`
	Value    interface{}     `json:"value"`
}

func NewBot(metadata Metadata, setting []SettingItem, rules []string) (*Bot, error) {
	cfg := &Config{}
	b := &Bot{
		Metadata: metadata,
		Setting:  setting,
		config:   cfg,
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