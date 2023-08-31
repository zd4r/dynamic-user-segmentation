package experiment

import (
	"context"

	"github.com/zd4r/dynamic-user-segmentation/internal/errs"
	experimentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/experiment"
)

func (s *Service) Delete(ctx context.Context, experiment *experimentModel.Experiment) error {
	rowsAffected, err := s.experimentRepository.Delete(ctx, experiment)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (s *Service) DeleteBatch(ctx context.Context, experiments []*experimentModel.Experiment) error {
	if len(experiments) != 0 {
		rowsAffected, err := s.experimentRepository.DeleteBatch(ctx, experiments)
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return errs.ErrNotFound
		}
	}

	return nil
}
