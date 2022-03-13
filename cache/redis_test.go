package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/subscan-explorer/subscan-common/util/xtime"
)

var (
	client *redis.Client
	cfg    *Config
	ctx    = context.Background()
)

func init() {
	cfg = &Config{
		Addr:         "127.0.0.1:6379",
		Username:     "",
		Password:     "",
		PoolSize:     1,
		MinIdleConns: 10,
		Dial:         xtime.Duration(100 * time.Millisecond),
		KeepAlive:    xtime.Duration(time.Second),
	}
	client = NewRedisClient(cfg)
}

func TestGet(t *testing.T) {
	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("key", val)
}
