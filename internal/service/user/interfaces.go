package user

import (
	"context"

	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"
	userRepo "github.com/zd4r/dynamic-user-segmentation/internal/repository/user"
)

type userRepository interface {
	Create(ctx context.Context, user *userModel.User) error
	Delete(ctx context.Context, user *userModel.User) (int64, error)
	GetSegments(ctx context.Context, user *userModel.User) ([]segmentModel.Segment, error)
}

var _ userRepository = (*userRepo.Repository)(nil)
