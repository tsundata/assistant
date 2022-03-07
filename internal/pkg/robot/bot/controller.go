package bot

import "github.com/tsundata/assistant/internal/pkg/robot/component"

type Controller struct {
	Instance    *Bot
	Config      *Config
	PluginParam map[string][]interface{}
	Comp        component.Component
}

func MockController() *Controller {
	return &Controller{}
}
