package orm

import "gorm.io/gorm"

func init() {
	cfg := &Config{
		DriverName: "mysql",
		DSN:        "root:@tcp(127.0.0.1:3306)/subscan_test?timeout=5s&readTimeout=5s", // change your test db instance .
	}
	_ = NewORM(cfg, &gorm.Config{})
}
