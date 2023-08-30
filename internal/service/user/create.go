package user

import (
	"context"

	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"
)

func (s *Service) Create(ctx context.Context, user *userModel.User) error {
	return s.userRepository.Create(ctx, user)
}
