package experiment

import (
	"context"

	experimentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/experiment"
)

func (s *Service) Create(ctx context.Context, segment *experimentModel.Experiment) error {
	return s.experimentRepository.Create(ctx, segment)
}
