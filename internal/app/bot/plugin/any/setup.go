package any

import (
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

func init() {
	bot.RegisterPlugin("any", setup)
}

func setup(c *bot.Controller) error {
	a := Any{}
	bot.GetConfig(c).AddPlugin(func(next bot.PluginHandler) bot.PluginHandler {
		a.Next = next
		return a
	})
	return nil
}
