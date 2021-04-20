module github.com/tsundata/assistant

go 1.16

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/PuerkitoBio/goquery v1.6.1
	github.com/RichardKnop/machinery/v2 v2.0.10
	github.com/go-ping/ping v0.0.0-20210216210419-25d1413fb7bb
	github.com/go-redis/redis/v8 v8.8.2
	github.com/go-resty/resty/v2 v2.6.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gofiber/fiber/v2 v2.8.0
	github.com/gogo/protobuf v1.3.2
	github.com/google/wire v0.5.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/influxdata/cron v0.0.0-20201006132531-4bb0a200dcbe
	github.com/influxdata/influxdb-client-go/v2 v2.2.3
	github.com/jmoiron/sqlx v1.3.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/rollbar/rollbar-go v1.3.0
	github.com/shirou/gopsutil/v3 v3.21.3
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/slack-go/slack v0.9.0
	github.com/sourcegraph/checkup v1.0.1-0.20200721114922-77e7567835d4
	github.com/spaolacci/murmur3 v1.1.0
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/tidwall/gjson v1.7.4
	github.com/uber/jaeger-client-go v2.26.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	github.com/valyala/fasthttp v1.23.0
	github.com/yuin/goldmark v1.3.5
	go.etcd.io/etcd v0.0.0-20201125193152-8a03d2e9614b
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.37.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1
