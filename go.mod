module github.com/tsundata/assistant

go 1.15

require (
	github.com/PuerkitoBio/goquery v1.6.0
	github.com/go-redis/redis/v8 v8.4.2
	github.com/gogo/protobuf v1.3.1
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/robertkrimen/otto v0.0.0-20200922221731-ef014fd054ac
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/slack-go/slack v0.7.2
	github.com/smallnest/rpcx v0.0.0-20201207055143-dff6bb8dd30b
	github.com/spf13/viper v1.7.1
	github.com/valyala/fasthttp v1.17.0
	go.etcd.io/etcd v0.0.0-20201125193152-8a03d2e9614b
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.34.0
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.5
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.0
