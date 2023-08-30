package segment

import (
	"context"

	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
)

func (s *Service) Create(ctx context.Context, segment *segmentModel.Segment) error {
	return s.segmentRepository.Create(ctx, segment)
}
