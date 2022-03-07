package bot

type Controller struct {
	Instance    *Bot
	Config      *Config
	PluginParam map[string][]interface{}
}

func MockController() *Controller {
	return &Controller{}
}
