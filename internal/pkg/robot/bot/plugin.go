package bot

import (
	"context"
	"log"
)

var (
	plugins = make(map[string]Plugin)
)

type SetupFunc func(c *Controller) error

type SetupPlugin struct {
	Action SetupFunc
}

func SetupPlugins(pluginRules []PluginRule) ([]Plugin, [][]interface{}) {
	var plugin []Plugin
	var params [][]interface{}
	for _, rule := range pluginRules {
		if p, ok := plugins[rule.Name]; ok {
			params = append(params, rule.Param)
			plugin = append(plugin, p)
		}
	}
	return plugin, params
}

func RegisterPlugin(name string, plugin Plugin) {
	if name == "" {
		panic("plugin must have a name")
	}
	if _, ok := plugins[name]; ok {
		panic("plugin named " + name + " already registered")
	}
	plugins[name] = plugin
	log.Println("[robot] register plugin:", name)
}

// -----

type PluginValue struct {
	Value string
	Stack []interface{}
}

type Plugin interface {
	Run(ctx context.Context, ctrl *Controller, param []interface{}, input PluginValue) (PluginValue, error)
	Name() string
}
