package core_pgx_pool

import (
	"errors"
	"fmt"

	core_postgres_pool "github.com/glebateee/taskapp/internal/core/repository/postgres/pool"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgxRow struct {
	pgx.Row
}

func (r pgxRow) Scan(dest ...any) error {

	if err := r.Row.Scan(dest...); err != nil {
		return mapErrors(err)
	}
	return nil
}

func mapErrors(err error) error {
	const (
		pgxViolatesForeignKeyErrCode = "23503"
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return core_postgres_pool.ErrNoRows
	}

	if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
		if pgErr.Code == pgxViolatesForeignKeyErrCode {
			return fmt.Errorf(
				"%v: %w",
				err,
				core_postgres_pool.ErrViolatesForeignKey,
			)
		}
	}
	return fmt.Errorf(
		"%v: %w",
		err,
		core_postgres_pool.ErrUnknown,
	)
}

type pgxRows struct {
	pgx.Rows
}

type pgxCmdTag struct {
	pgconn.CommandTag
}
