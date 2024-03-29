package config

// === App ===

type SvcAddr struct {
	Id       string `json:"id" yaml:"id"`
	Dtm      string `json:"dtm" yaml:"dtm"`
	Chatbot  string `json:"chatbot" yaml:"chatbot"`
	Message  string `json:"message" yaml:"message"`
	Middle   string `json:"middle" yaml:"middle"`
	Workflow string `json:"workflow" yaml:"workflow"`
	User     string `json:"user" yaml:"user"`
	Nlp      string `json:"nlp" yaml:"nlp"`
	Storage  string `json:"storage" yaml:"storage"`
	Task     string `json:"task" yaml:"task"`
	Bot      string `json:"bot" yaml:"bot"`
}

// Http http config
type Http struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

// Rpc http config
type Rpc struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

// Web config
type Web struct {
	Url string `json:"url" yaml:"url"`
}

// Gateway config
type Gateway struct {
	Url string `json:"url" yaml:"url"`
}

// Storage config
type Storage struct {
	Adapter string `json:"adapter" yaml:"adapter"`
	Dir     string `json:"dir" yaml:"dir"`
}

// Jwt config
type Jwt struct {
	Secret string `json:"secret" yaml:"secret"`
}

// === Middleware ===

// Mysql config
type Mysql struct {
	Dsn string `json:"dsn" yaml:"dsn"`
}

// Redis config
type Redis struct {
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
}

// Influx config
type Influx struct {
	Token  string `json:"token" yaml:"token"`
	Org    string `json:"org" yaml:"org"`
	Bucket string `json:"bucket" yaml:"bucket"`
	Url    string `json:"url" yaml:"url"`
}

// Jaeger config
type Jaeger struct {
	Reporter struct {
		LocalAgentHostPort string `json:"localAgentHostPort" yaml:"localAgentHostPort"`
	} `json:"reporter" yaml:"reporter"`
	Sampler struct {
		Type  string  `json:"type" yaml:"type"`
		Param float64 `json:"param" yaml:"param"`
	} `json:"sampler" yaml:"sampler"`
}

// Nats config
type Nats struct {
	Url string `json:"url" yaml:"url"`
}

// Rabbitmq config
type Rabbitmq struct {
	Url string `json:"url" yaml:"url"`
}

// === Vendor ===

// Rollbar config
type Rollbar struct {
	Token       string `json:"token" yaml:"token"`
	Environment string `json:"environment" yaml:"environment"`
}

// Newrelic config
type Newrelic struct {
	Name    string `json:"name" yaml:"name"`
	License string `json:"license" yaml:"license"`
}
