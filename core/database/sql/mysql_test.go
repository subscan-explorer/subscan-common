package sql

import (
	"context"
	"time"

	"github.com/subscan-explorer/subscan-common/core/util/xtime"
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
		DSN:         "root:@tcp(127.0.0.1:3306)/subscan_test?timeout=5s&readTimeout=5s", // change your test db instance .
		Active:      10,
		Idle:        10,
		IdleTimeout: xtime.Duration(time.Second),
	}
	db = NewMySQL(cfg)
	_, _ = db.Exec(context.TODO(), `CREATE TABLE test_tbl (
  id int(11) unsigned NOT NULL AUTO_INCREMENT,
  title varchar(11) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  author varchar(11) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  date datetime DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`)
}
