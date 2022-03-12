package sql

import (
	"fmt"
	"time"

	"github.com/subscan-explorer/subscan-common/util/xtime"
)

// CREATE TABLE IF NOT EXISTS `test_tbl`(
// 	`id` INT UNSIGNED AUTO_INCREMENT,
// 	`title` VARCHAR(100) NOT NULL,
// 	`author` VARCHAR(40) NOT NULL,
// 	`date` DATE,
// 	PRIMARY KEY ( `id` )
//  )ENGINE=InnoDB DEFAULT CHARSET=utf8;

// INSERT INTO test_tbl (title, author, date)VALUES("mysql test", "test", NOW());

var db *DB

func init() {
	cfg := &Config{
		DSN:         "root:123456@tcp(192.168.50.144:3306)/test?timeout=5s&readTimeout=5s", // change your test db instance .
		Active:      10,
		Idle:        10,
		IdleTimeout: xtime.Duration(time.Second),
	}
	db = NewMySQL(cfg)
	fmt.Println(db)
}
