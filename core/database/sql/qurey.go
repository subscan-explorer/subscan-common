package sql

import (
	"context"
	"database/sql"
	xsql "database/sql"
)

// Exec is wrap mysql ExecContext
func (db *DB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.write.ExecContext(ctx, query, args...)
}

// QureyRow is wrap mysql QueryRowContext
func (db *DB) QureyRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	idx := db.readIndex()
	if len(db.read) > idx {
		return db.read[(idx)%len(db.read)].QueryRowContext(ctx, query, args...)
	}
	return db.write.QueryRowContext(ctx, query, args...)
}

// Qurey is wrap mysql QueryContext
func (db *DB) Qurey(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	idx := db.readIndex()
	if len(db.read) > idx {
		if rows, err = db.read[(idx)%len(db.read)].QueryContext(ctx, query, args...); err != nil {
			return
		}
	}
	return db.write.QueryContext(ctx, query, args...)
}

// Begin is wrap mysql BeginTx
func (db *DB) Begin(ctx context.Context, opts *xsql.TxOptions) (*xsql.Tx, error) {
	return db.write.BeginTx(ctx, opts)
}
