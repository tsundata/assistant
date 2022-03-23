package bot

type Config struct {
	// Plugin stack
	Plugin []Plugin
	// Compiled plugin stack
	PluginChain PluginHandler
	// registry plugin
	registry map[string]PluginHandler
}

func (c *Config) AddPlugin(m Plugin) {
	c.Plugin = append(c.Plugin, m)
}

func (c *Config) RegisterHandler(h PluginHandler) {
	if c.registry == nil {
		c.registry = make(map[string]PluginHandler)
	}
	c.registry[h.Name()] = h
}

func (c *Config) Handler(name string) PluginHandler {
	if c.registry == nil {
		return nil
	}
	if h, ok := c.registry[name]; ok {
		return h
	}
	return nil
}

func (c *Config) Handlers() []PluginHandler {
	if c.registry == nil {
		return nil
	}
	hs := make([]PluginHandler, 0, len(c.registry))
	for k := range c.registry {
		hs = append(hs, c.registry[k])
	}
	return hs
}

func GetConfig(c *Controller) *Config {
	return c.Config
}
