package bot

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
)

var (
	plugins = make(map[string]SetupPlugin)
)

type SetupFunc func(c *Controller) error

type SetupPlugin struct {
	Action SetupFunc
}

func ListPlugins() map[string][]string {
	p := make(map[string][]string)
	for s := range plugins {
		p["bot"] = append(p["bot"], s)
	}
	return p
}

func SetupPlugins(c *Controller, pluginRules []PluginRule) error {
	pluginRules = append(pluginRules, PluginRule{Name: "end"})
	for _, rule := range pluginRules {
		if plugin, ok := plugins[rule.Name]; ok {
			err := plugin.Action(c)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("plugin/%s", rule.Name))
			}
		}
	}
	return nil
}

func RegisterPlugin(name string, action SetupFunc) {
	registerPlugin(name, SetupPlugin{
		action,
	})
}

func registerPlugin(name string, plugin SetupPlugin) {
	if name == "" {
		panic("plugin must have a name")
	}
	if _, ok := plugins[name]; ok {
		panic("plugin named " + name + " already registered")
	}
	plugins[name] = plugin
}

// -----

type (
	Plugin func(next PluginHandler) PluginHandler

	PluginHandler interface {
		Run(ctx context.Context, input interface{}) (interface{}, error)
		Name() string
	}

	HandlerFunc func(ctx context.Context, input interface{}) (interface{}, error)
)

func (f HandlerFunc) Run(ctx context.Context, input interface{}) (interface{}, error) {
	return f(ctx, input)
}

func (f HandlerFunc) Name() string {
	return "handlerfunc"
}

// Error returns err with 'plugin/name: ' prefixed to it.
func Error(name string, err error) error {
	return fmt.Errorf("%s/%s: %s", "plugin", name, err)
}

func NextOrFailure(name string, next PluginHandler, ctx context.Context, input interface{}) (interface{}, error) {
	if next != nil {
		return next.Run(ctx, input)
	}
	return nil, Error(name, errors.New("no next plugin found"))
}
