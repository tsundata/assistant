module github.com/tsundata/assistant

go 1.16

require (
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/PuerkitoBio/goquery v1.6.1
	github.com/go-redis/redis/v8 v8.6.0
	github.com/go-resty/resty/v2 v2.5.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/influxdata/cron v0.0.0-20201006132531-4bb0a200dcbe
	github.com/influxdata/influxdb-client-go/v2 v2.2.2
	github.com/jmoiron/sqlx v1.3.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/shirou/gopsutil/v3 v3.21.2
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/slack-go/slack v0.8.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tidwall/gjson v1.6.8
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	github.com/valyala/fasthttp v1.21.0
	github.com/yuin/goldmark v1.3.2
	go.etcd.io/etcd v0.0.0-20201125193152-8a03d2e9614b
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.35.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1
