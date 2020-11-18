module github.com/tsundata/assistant

go 1.15

require (
	github.com/robertkrimen/otto v0.0.0-20200922221731-ef014fd054ac
	github.com/robfig/cron/v3 v3.0.1
	github.com/slack-go/slack v0.7.2
	github.com/smallnest/rpcx v0.0.0-20201027145221-c31b15be63d4
	github.com/spf13/viper v1.7.1
	github.com/valyala/fasthttp v1.17.0
	go.uber.org/zap v1.16.0
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.5
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.0
