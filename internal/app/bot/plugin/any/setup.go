package any

import "github.com/tsundata/assistant/internal/pkg/robot/plugin"

func init() {
	plugin.Register("any", setup)
}

func setup(c *plugin.Controller) error {
	a := Any{}
	plugin.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		a.Next = next
		return a
	})
	return nil
}
