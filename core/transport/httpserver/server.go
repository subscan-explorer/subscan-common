package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/subscan-explorer/subscan-common/core/log"
	"github.com/subscan-explorer/subscan-common/core/util/xtime"
)

type Config struct {
	Addr         string
	Handler      http.Handler
	ReadTimeout  xtime.Duration
	WriteTimeout xtime.Duration
}

var httpServer *http.Server

func Run(c *Config) {
	httpServer = &http.Server{
		Addr:         c.Addr,
		Handler:      c.Handler,
		ReadTimeout:  time.Duration(c.ReadTimeout),
		WriteTimeout: time.Duration(c.WriteTimeout),
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func Shutdown(ctx context.Context) (err error) {
	if err = httpServer.Shutdown(ctx); err != nil {
		log.Errorf("http server stop error %v", err)
	}
	return
}
