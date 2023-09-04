package segment

import (
	"context"

	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"
)

func (s *Service) GetBySlug(ctx context.Context, slug string) (*segmentModel.Segment, error) {
	return s.segmentRepository.GetBySlug(ctx, slug)
}

func (s *Service) GetUsersBySlug(ctx context.Context, slug string) ([]userModel.User, error) {
	return s.segmentRepository.GetUsersBySlug(ctx, slug)
}
