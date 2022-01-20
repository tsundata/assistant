package plugin

type Config struct {
	// Plugin stack
	Plugin []Plugin
	// Compiled plugin stack
	PluginChain Handler
	// registry plugin
	registry map[string]Handler
}

func (c *Config) AddPlugin(m Plugin) {
	c.Plugin = append(c.Plugin, m)
}

func (c *Config) RegisterHandler(h Handler) {
	if c.registry == nil {
		c.registry = make(map[string]Handler)
	}
	c.registry[h.Name()] = h
}

func (c *Config) Handler(name string) Handler {
	if c.registry == nil {
		return nil
	}
	if h, ok := c.registry[name]; ok {
		return h
	}
	return nil
}

func (c *Config) Handlers() []Handler {
	if c.registry == nil {
		return nil
	}
	hs := make([]Handler, 0, len(c.registry))
	for k := range c.registry {
		hs = append(hs, c.registry[k])
	}
	return hs
}

func GetConfig(c *Controller) *Config {
	return c.Config
}
