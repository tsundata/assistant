module github.com/tsundata/assistant

go 1.15

require (
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/PuerkitoBio/goquery v1.6.1
	github.com/go-redis/redis/v8 v8.4.2
	github.com/gogo/protobuf v1.3.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/influxdata/cron v0.0.0-20201006132531-4bb0a200dcbe
	github.com/influxdata/influxdb-client-go/v2 v2.2.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/robertkrimen/otto v0.0.0-20200922221731-ef014fd054ac
	github.com/shirou/gopsutil/v3 v3.20.12
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/slack-go/slack v0.7.2
	github.com/spf13/viper v1.7.1
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	github.com/valyala/fasthttp v1.17.0
	go.etcd.io/bbolt v1.3.5
	go.etcd.io/etcd v0.0.0-20201125193152-8a03d2e9614b
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.34.0
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.0
