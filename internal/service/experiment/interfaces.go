package experiment

import (
	"context"

	experimentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/experiment"
	experimentRepo "github.com/zd4r/dynamic-user-segmentation/internal/repository/experiment"
)

type experimentRepository interface {
	Create(ctx context.Context, segment *experimentModel.Experiment) error
	Delete(ctx context.Context, segment *experimentModel.Experiment) (int64, error)
}

var _ experimentRepository = (*experimentRepo.Repository)(nil)
