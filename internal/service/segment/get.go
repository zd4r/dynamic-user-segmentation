package segment

import (
	"context"

	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
)

func (s *Service) GetBySlug(ctx context.Context, slug string) (*segmentModel.Segment, error) {
	return s.segmentRepository.GetBySlug(ctx, slug)
}
