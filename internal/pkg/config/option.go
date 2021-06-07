package config

// === App ===

// Http http config
type Http struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Mode string `json:"mode"`
}

// Rpc http config
type Rpc struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// Web config
type Web struct {
	Url string `json:"url"`
}

// Gateway config
type Gateway struct {
	Url string `json:"url"`
}

// Plugin spider config
type Plugin struct {
	Path string `json:"path"`
}

// Storage config
type Storage struct {
	Path string `json:"path"`
}

// === Middleware ===

// Mysql config
type Mysql struct {
	Url string `json:"url"`
}

// Redis config
type Redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
}

// Etcd config
type Etcd struct {
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Influx config
type Influx struct {
	Token  string `json:"token"`
	Org    string `json:"org"`
	Bucket string `json:"bucket"`
	Url    string `json:"url"`
}

// Rabbitmq config
type Rabbitmq struct {
	Url string `json:"url"`
}

// Jaeger config
type Jaeger struct {
	ServiceName string `json:"serviceName"`
	Reporter    struct {
		LocalAgentHostPort string `json:"localAgentHostPort"`
	} `json:"reporter"`
	Sampler struct {
		Type  string  `json:"type"`
		Param float64 `json:"param"`
	} `json:"sampler"`
}

// Nats config
type Nats struct {
	Url string `json:"url"`
}

// === Vendor ===

// Slack config
type Slack struct {
	Verification string `json:"verification"`
	Signing      string `json:"signing"`
	Token        string `json:"token"`
	Webhook      string `json:"webhook"`
}

// Rollbar config
type Rollbar struct {
	Token       string `json:"token"`
	Environment string `json:"environment"`
}

// Telegram config
type Telegram struct {
	Token string `json:"token"`
}
