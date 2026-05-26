package users_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/glebateee/taskapp/internal/core/domain"
	core_errors "github.com/glebateee/taskapp/internal/core/errors"
	core_postgres_pool "github.com/glebateee/taskapp/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) GetUser(
	ctx context.Context,
	userID int,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT 
				id,
				version,
				full_name,
				phone_number
			  FROM taskapp.users
			  WHERE id = $1;`
	row := r.pool.QueryRow(ctx, query, userID)
	var userModel UserModel
	if err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id='%d': %w", userID, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("scan user: %w", err)
	}
	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.PhoneNumber,
	)
	return userDomain, nil
}
