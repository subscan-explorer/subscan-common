package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	xredis "github.com/go-redis/redis/v8"
	"github.com/subscan-explorer/subscan-common/core/util/xtime"
)

var (
	client *xredis.Client
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

// go test -v -test.run TestGet
func TestGet(t *testing.T) {
	val, err := client.Get(ctx, "hello").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("value", val)
}
