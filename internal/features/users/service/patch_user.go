package users_service

import (
	"context"
	"fmt"

	"github.com/glebateee/taskapp/internal/core/domain"
)

func (s UsersService) PatchUser(
	ctx context.Context,
	userID int,
	patch domain.UserPatch,
) (domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}
	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("apply user patch: %w", err)
	}
	user, err = s.usersRepository.PatchUser(ctx, userID, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("patch user: %w", err)
	}

	return user, nil
}
