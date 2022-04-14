package bot

import "github.com/tsundata/assistant/internal/pkg/robot/component"

type Controller struct {
	Instance *Bot
	Config   *Config
	Comp     component.Component

	PluginParam map[string][]interface{}
}

func MockController(param ...map[string][]interface{}) *Controller {
	pluginParam := make(map[string][]interface{})
	for _, p := range param {
		for k, v := range p {
			pluginParam[k] = v
		}
	}
	return &Controller{PluginParam: pluginParam}
}
