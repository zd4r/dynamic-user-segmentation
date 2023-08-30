package segment

import (
	"context"

	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
	segmentRepo "github.com/zd4r/dynamic-user-segmentation/internal/repository/segment"
)

type segmentRepository interface {
	Create(ctx context.Context, segment *segmentModel.Segment) error
	Delete(ctx context.Context, segment *segmentModel.Segment) (int64, error)
}

var _ segmentRepository = (*segmentRepo.Repository)(nil)
