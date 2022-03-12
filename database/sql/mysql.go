package sql

import (
	"github.com/subscan-explorer/subscan-common/log"
	"github.com/subscan-explorer/subscan-common/util/xtime"
)

type Config struct {
	DSN         string         // write data source name.
	ReadDSN     []string       // read data source name.
	Active      int            // pool
	Idle        int            // pool
	IdleTimeout xtime.Duration // connect max life time.
}

// NewMySQL new db instance .
func NewMySQL(c *Config) (db *DB) {
	db, err := Open(c)
	if err != nil {
		log.Error("open mysql error(%v)", err)
		panic(err)
	}
	return
}
