package user

import (
	"context"

	"github.com/zd4r/dynamic-user-segmentation/internal/errs"
)

func (s *Service) Delete(ctx context.Context, userId int) error {
	rowsAffected, err := s.userRepository.Delete(ctx, userId)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}
