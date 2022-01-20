package bot

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/robot/plugin"
)

type Bot struct {
	*Metadata
	Setting []SettingItem
	plugin  []plugin.Plugin
	config  *plugin.Config
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

func NewBot(rules []string) (*Bot, error) {
	c := &plugin.Controller{
		Config: &plugin.Config{},
	}
	b := Bot{
		config: c.Config,
	}

	// setup plugins
	err := plugin.SetupPlugins(c, rules)
	if err != nil {
		return nil, err
	}
	b.plugin = c.Config.Plugin

	// plugin chain
	var stack plugin.Handler
	for i := len(b.plugin) - 1; i >= 0; i-- {
		stack = b.plugin[i](stack)
		fmt.Println(b.plugin)
		b.config.RegisterHandler(stack)
	}
	b.config.PluginChain = stack

	return &b, nil
}

func (b *Bot) Run(ctx context.Context) error {
	_, err := b.config.PluginChain.Run(ctx, nil)
	return err
}
