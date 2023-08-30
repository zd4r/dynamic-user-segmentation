package v1

import (
	"context"

	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"
	segmentSrv "github.com/zd4r/dynamic-user-segmentation/internal/service/segment"
	userSrv "github.com/zd4r/dynamic-user-segmentation/internal/service/user"
)

type userService interface {
	Delete(ctx context.Context, user *userModel.User) error
	Create(ctx context.Context, user *userModel.User) error
}

var _ userService = (*userSrv.Service)(nil)

type segmentService interface {
	Delete(ctx context.Context, segment *segmentModel.Segment) error
	Create(ctx context.Context, segment *segmentModel.Segment) error
}

var _ segmentService = (*segmentSrv.Service)(nil)
