module github.com/tsundata/assistant

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.3
	github.com/robfig/cron/v3 v3.0.1
	github.com/smallnest/rpcx v0.0.0-20201027145221-c31b15be63d4
	github.com/spf13/viper v1.7.1
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.5
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.0
