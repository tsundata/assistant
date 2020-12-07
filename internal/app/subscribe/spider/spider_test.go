package spider

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"testing"
	"time"
)

func TestSpider(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
	})
	spider := New(rdb, viper.New())
	spider.Cron()

	time.Sleep(100 * time.Minute)
}
