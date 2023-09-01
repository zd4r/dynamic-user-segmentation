package segment

import (
	"context"

	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
	segmentRepo "github.com/zd4r/dynamic-user-segmentation/internal/repository/segment"
)

type segmentRepository interface {
	Create(ctx context.Context, slug string) error
	DeleteBySlug(ctx context.Context, slug string) (int64, error)
	GetBySlug(ctx context.Context, slug string) (*segmentModel.Segment, error)
}

var _ segmentRepository = (*segmentRepo.Repository)(nil)
