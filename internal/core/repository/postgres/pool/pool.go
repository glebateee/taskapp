package core_postgres_pool

import (
	"context"
	"time"
)

type Pool interface {
	Close()
	Exec(ctx context.Context, sql string, arguments ...any) (CmdTag, error)
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row
	OpTimeout() time.Duration
}

type CmdTag interface {
	RowsAffected() int64
}

type Rows interface {
	Close()
	Err() error
	Next() bool
	Scan(dest ...any) error
}

type Row interface {
	Scan(dest ...any) error
}
