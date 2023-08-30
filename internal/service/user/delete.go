package user

import (
	"context"

	"github.com/zd4r/dynamic-user-segmentation/internal/errs"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"
)

func (s *Service) Delete(ctx context.Context, user *userModel.User) error {
	rowsAffected, err := s.userRepository.Delete(ctx, user)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}
