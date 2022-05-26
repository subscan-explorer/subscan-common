package orm

import (
	"github.com/subscan-explorer/subscan-common/core/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DriverName string
	DSN        string // DSN data source name
}

func NewORM(c *Config, gormConf *gorm.Config) (db *gorm.DB) {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DriverName:                c.DriverName,
		DSN:                       c.DSN, // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: true,  // 根据当前 MySQL 版本自动配置
	}), gormConf)
	if err != nil {
		log.Fatalf("gorm.Open dsn %s error %v", c.DSN, err)
	}
	return
}
