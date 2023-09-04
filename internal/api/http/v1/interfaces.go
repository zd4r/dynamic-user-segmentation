package v1

import (
	"context"
	"time"

	experimentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/experiment"
	reportModel "github.com/zd4r/dynamic-user-segmentation/internal/model/report"
	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"
	reportSrv "github.com/zd4r/dynamic-user-segmentation/internal/service/report"

	experimentSrv "github.com/zd4r/dynamic-user-segmentation/internal/service/experiment"
	segmentSrv "github.com/zd4r/dynamic-user-segmentation/internal/service/segment"
	userSrv "github.com/zd4r/dynamic-user-segmentation/internal/service/user"
)

type userService interface {
	Create(ctx context.Context, user *userModel.User) error
	GetPercentOfAllUsers(ctx context.Context, percent int) ([]userModel.User, error)
	Delete(ctx context.Context, userId int) error
	GetSegments(ctx context.Context, userId int) ([]segmentModel.Segment, error)
}

var _ userService = (*userSrv.Service)(nil)

type segmentService interface {
	Create(ctx context.Context, slug string) error
	DeleteBySlug(ctx context.Context, slug string) error
	GetBySlug(ctx context.Context, slug string) (*segmentModel.Segment, error)
	GetUsersBySlug(ctx context.Context, slug string) ([]userModel.User, error)
}

var _ segmentService = (*segmentSrv.Service)(nil)

type experimentService interface {
	CreateBatch(ctx context.Context, experiments []experimentModel.Experiment) error
	DeleteBatch(ctx context.Context, experiments []experimentModel.Experiment) error
}

var _ experimentService = (*experimentSrv.Service)(nil)

type reportService interface {
	CreateBatchRecord(ctx context.Context, records []reportModel.Record) error
	GetRecordsInIntervalByUser(ctx context.Context, userId int, from, to time.Time) ([]reportModel.Record, error)
}

var _ reportService = (*reportSrv.Service)(nil)
