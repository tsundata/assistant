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

func SetupPlugins(c *Controller, pluginRules []PluginRule) error {
	pluginRules = append(pluginRules, PluginRule{Name: "end"})
	for _, rule := range pluginRules {
		if plugin, ok := plugins[rule.Name]; ok {
			c.PluginParam[rule.Name] = rule.Param
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
	fmt.Println("[robot] register plugin", name)
}

// -----

type (
	Plugin func(next PluginHandler) PluginHandler

	PluginHandler interface {
		Run(ctx context.Context, ctrl *Controller, input interface{}) (interface{}, error)
		Name() string
	}
)

// Error returns err with 'plugin/name: ' prefixed to it.
func Error(name string, err error) error {
	return fmt.Errorf("%s/%s: %s", "plugin", name, err)
}

func NextOrFailure(ctx context.Context, name string, next PluginHandler, ctrl *Controller, input interface{}) (interface{}, error) {
	if next != nil {
		return next.Run(ctx, ctrl, input)
	}
	return nil, Error(name, errors.New("no next plugin found"))
}

func Param(ctrl *Controller, p PluginHandler) interface{} {
	return ctrl.PluginParam[p.Name()]
}
