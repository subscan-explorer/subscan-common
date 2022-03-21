package orm

import (
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	cfg := &Config{
		DriverName: "mysql",
		DSN:        "root:123456@tcp(192.168.50.144:3306)/test?timeout=5s&readTimeout=5s", // change your test db instance .
	}
	db = NewORM(cfg)
	fmt.Println(db)
}
