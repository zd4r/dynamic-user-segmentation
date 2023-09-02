package experiment

import (
	"context"

	experimentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/experiment"
	experimentRepo "github.com/zd4r/dynamic-user-segmentation/internal/repository/experiment"
)

//go:generate mockery --name experimentRepository --structname ExperimentRepository
type experimentRepository interface {
	Create(ctx context.Context, segment *experimentModel.Experiment) error
	CreateBatch(ctx context.Context, experiments []experimentModel.Experiment) error
	Delete(ctx context.Context, segment *experimentModel.Experiment) (int64, error)
	DeleteBatch(ctx context.Context, experiments []experimentModel.Experiment) (int64, error)
	DeleteAllExpired(ctx context.Context) (int64, error)
}

var _ experimentRepository = (*experimentRepo.Repository)(nil)
