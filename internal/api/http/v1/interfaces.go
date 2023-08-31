package v1

import (
	"context"

	experimentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/experiment"
	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"

	experimentSrv "github.com/zd4r/dynamic-user-segmentation/internal/service/experiment"
	segmentSrv "github.com/zd4r/dynamic-user-segmentation/internal/service/segment"
	userSrv "github.com/zd4r/dynamic-user-segmentation/internal/service/user"
)

type userService interface {
	Create(ctx context.Context, user *userModel.User) error
	Delete(ctx context.Context, userId int) error
	GetSegments(ctx context.Context, userId int) ([]segmentModel.Segment, error)
}

var _ userService = (*userSrv.Service)(nil)

type segmentService interface {
	Create(ctx context.Context, segment *segmentModel.Segment) error
	DeleteBySlug(ctx context.Context, slug string) error
	GetBySlug(ctx context.Context, slug string) (*segmentModel.Segment, error)
}

var _ segmentService = (*segmentSrv.Service)(nil)

type experimentService interface {
	CreateBatch(ctx context.Context, experiments []experimentModel.Experiment) error
	DeleteBatch(ctx context.Context, experiments []experimentModel.Experiment) error
}

var _ experimentService = (*experimentSrv.Service)(nil)
