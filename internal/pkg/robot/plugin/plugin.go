package plugin

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
)

var (
	plugins = make(map[string]BotPlugin)
)

type SetupFunc func(c *Controller) error

type BotPlugin struct {
	Action SetupFunc
}

func RegisterPlugin(name string, plugin BotPlugin) {
	if name == "" {
		panic("plugin must have a name")
	}
	if _, ok := plugins[name]; ok {
		panic("plugin named " + name + " already registered")
	}
	plugins[name] = plugin
}

func ListPlugins() map[string][]string {
	p := make(map[string][]string)
	for s := range plugins {
		p["bot"] = append(p["bot"], s)
	}
	return p
}

func SetupPlugins(c *Controller, pluginRules []string) error {
	pluginRules = append(pluginRules, "end")
	for _, name := range pluginRules {
		if plugin, ok := plugins[name]; ok {
			err := plugin.Action(c)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("plugin/%s", name))
			}
		}
	}
	return nil
}

func Register(name string, action SetupFunc) {
	RegisterPlugin(name, BotPlugin{
		action,
	})
}

// -----

type (
	Plugin func(next Handler) Handler

	Handler interface {
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

func NextOrFailure(name string, next Handler, ctx context.Context, input interface{}) (interface{}, error) {
	if next != nil {
		return next.Run(ctx, input)
	}
	return nil, Error(name, errors.New("no next plugin found"))
}
