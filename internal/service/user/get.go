package user

import (
	"context"

	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
)

func (s *Service) GetSegments(ctx context.Context, userId int) ([]segmentModel.Segment, error) {
	return s.userRepository.GetSegments(ctx, userId)
}
