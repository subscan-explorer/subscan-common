package redis

import (
	"net"
	"time"

	xredis "github.com/go-redis/redis/v8"
	"github.com/subscan-explorer/subscan-common/util/xtime"
)

// Config is redis client config .
type Config struct {
	Addr         string
	Username     string
	Password     string
	PoolSize     int
	MinIdleConns int
	Dial         xtime.Duration
	KeepAlive    xtime.Duration
}

// NewRedisClient initial redis client .
func NewRedisClient(c *Config) *xredis.Client {
	dialer := &net.Dialer{
		Timeout:   time.Duration(c.Dial),
		KeepAlive: time.Duration(c.KeepAlive),
	}
	return xredis.NewClient(&xredis.Options{
		Addr:         c.Addr,
		Dialer:       dialer.DialContext,
		Username:     c.Username,
		Password:     c.Password,
		PoolSize:     c.PoolSize,
		MinIdleConns: c.MinIdleConns,
	})
}
