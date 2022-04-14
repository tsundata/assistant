package keyword

import (
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

func init() {
	bot.RegisterPlugin("keyword", setup)
}

func setup(c *bot.Controller) error {
	a := Keyword{}
	bot.GetConfig(c).AddPlugin(func(next bot.PluginHandler) bot.PluginHandler {
		a.Next = next
		return a
	})
	return nil
}
