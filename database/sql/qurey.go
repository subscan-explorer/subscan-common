package sql

import (
	"context"
	"database/sql"
)

// QureyRow is wrap mysql qureyrow
func (db *DB) QureyRow(ctx context.Context, query string, args ...interface{}) (row *sql.Row) {
	idx := db.readIndex()
	if len(db.read) > idx {
		row = db.read[(idx)%len(db.read)].QueryRowContext(ctx, query, args...)
	}
	row = db.write.QueryRowContext(ctx, query, args...)
	return
}

// Qurey is wrap mysql qurey
func (db *DB) Qurey(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	idx := db.readIndex()
	if len(db.read) > idx {
		if rows, err = db.read[(idx)%len(db.read)].QueryContext(ctx, query, args...); err != nil {
			return
		}
	}
	return db.write.QueryContext(ctx, query, args...)
}
