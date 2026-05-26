package users_repository_postgres

import (
	"context"
	"fmt"

	core_errors "github.com/glebateee/taskapp/internal/core/errors"
)

func (r *UsersRepository) DeleteUser(
	ctx context.Context,
	userID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `DELETE
			  FROM taskapp.users
			  WHERE id = $1;`
	tag, err := r.pool.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("user with id='%d: %w", userID, core_errors.ErrNotFound)
	}
	return nil
}
