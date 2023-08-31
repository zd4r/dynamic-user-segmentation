package experiment

import (
	"context"

	"github.com/zd4r/dynamic-user-segmentation/internal/errs"
	experimentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/experiment"
)

func (s *Service) Delete(ctx context.Context, segment *experimentModel.Experiment) error {
	rowsAffected, err := s.experimentRepository.Delete(ctx, segment)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}
