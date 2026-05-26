package core_pgx_pool

import (
	"errors"

	core_postgres_pool "github.com/glebateee/taskapp/internal/core/repository/postgres/pool"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgxRow struct {
	pgx.Row
}

func (r pgxRow) Scan(dest ...any) error {
	if err := r.Row.Scan(dest...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return core_postgres_pool.ErrNoRows
		}
		return err
	}
	return nil
}

type pgxRows struct {
	pgx.Rows
}

type pgxCmdTag struct {
	pgconn.CommandTag
}
