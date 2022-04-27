package bot

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/util"
	"log"
)

var (
	plugins = make(map[string]Plugin)
)

type SetupFunc func(c *Controller) error

type SetupPlugin struct {
	Action SetupFunc
}

func SetupPlugins(pluginRules []PluginRule) ([]Plugin, [][]util.Value) {
	var plugin []Plugin
	var params [][]util.Value
	for _, rule := range pluginRules {
		if p, ok := plugins[rule.Name]; ok {
			var tmp []util.Value
			for _, item := range rule.Param {
				tmp = append(tmp, util.Variable(item))
			}
			params = append(params, tmp)
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
	Run(ctx context.Context, ctrl *Controller, param []util.Value, input PluginValue) (PluginValue, error)
	Name() string
}
