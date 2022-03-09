package transport

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/subscan-explorer/subscan-common/util/xtime"
)

// go test -v -test.run TestRun
func TestRun(t *testing.T) {
	cfg := &Config{
		Addr:         ":8080",
		WriteTimeout: xtime.Duration(time.Second),
		ReadTimeout:  xtime.Duration(time.Second),
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.GET("/health", health)
	cfg.Handler = r
	Run(cfg)
	time.Sleep(time.Minute * 1)
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
