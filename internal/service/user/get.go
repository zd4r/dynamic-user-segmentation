package user

import (
	"context"

	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"
)

func (s *Service) GetSegments(ctx context.Context, user *userModel.User) ([]segmentModel.Segment, error) {
	return s.userRepository.GetSegments(ctx, user)
}
