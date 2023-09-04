package segment

import (
	"context"

	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"
	segmentRepo "github.com/zd4r/dynamic-user-segmentation/internal/repository/segment"
)

//go:generate mockery --name segmentRepository --structname SegmentRepository
type segmentRepository interface {
	Create(ctx context.Context, slug string) error
	DeleteBySlug(ctx context.Context, slug string) (int64, error)
	GetBySlug(ctx context.Context, slug string) (*segmentModel.Segment, error)
	GetUsersBySlug(ctx context.Context, slug string) ([]userModel.User, error)
}

var _ segmentRepository = (*segmentRepo.Repository)(nil)
