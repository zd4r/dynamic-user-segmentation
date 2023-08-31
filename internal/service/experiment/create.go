package experiment

import (
	"context"

	experimentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/experiment"
)

func (s *Service) Create(ctx context.Context, experiment *experimentModel.Experiment) error {
	return s.experimentRepository.Create(ctx, experiment)
}

func (s *Service) CreateBatch(ctx context.Context, experiments []*experimentModel.Experiment) error {
	if len(experiments) != 0 {
		return s.experimentRepository.CreateBatch(ctx, experiments)
	}
	return nil
}
