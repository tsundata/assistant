package filter

import "github.com/tsundata/assistant/internal/pkg/robot/plugin"

func init() {
	plugin.Register("filter", setup)
}

func setup(c *plugin.Controller) error {
	a := Filter{}
	plugin.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		a.Next = next
		return a
	})
	return nil
}
