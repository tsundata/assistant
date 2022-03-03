package bot

import "context"

type Controller struct {
	Ctx         context.Context
	Instance    *Bot
	Config      *Config
	PluginParam map[string][]interface{}
}

func MockController() *Controller {
	return &Controller{}
}
