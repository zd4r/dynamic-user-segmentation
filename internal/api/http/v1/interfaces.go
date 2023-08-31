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
	Delete(ctx context.Context, user *userModel.User) error
	Create(ctx context.Context, user *userModel.User) error
	GetSegments(ctx context.Context, user *userModel.User) ([]segmentModel.Segment, error)
}

var _ userService = (*userSrv.Service)(nil)

type segmentService interface {
	Delete(ctx context.Context, segment *segmentModel.Segment) error
	Create(ctx context.Context, segment *segmentModel.Segment) error
}

var _ segmentService = (*segmentSrv.Service)(nil)

type experimentService interface {
	Delete(ctx context.Context, segment *experimentModel.Experiment) error
	Create(ctx context.Context, segment *experimentModel.Experiment) error
}

var _ experimentService = (*experimentSrv.Service)(nil)
