package end

import "github.com/tsundata/assistant/internal/pkg/robot/plugin"

func init() {
	plugin.Register("end", setup)
}

func setup(c *plugin.Controller) error {
	a := End{}
	plugin.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		a.Next = next
		return a
	})
	return nil
}
