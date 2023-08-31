package experiment

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/zd4r/dynamic-user-segmentation/internal/client/pg"
	experimentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/experiment"
)

const (
	experimentTableName = `experiment`
)

type Repository struct {
	client pg.Client
}

func NewRepository(client pg.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) Create(ctx context.Context, experiment *experimentModel.Experiment) error {
	builder := sq.Insert(experimentTableName).
		PlaceholderFormat(sq.Dollar).
		Columns("user_id", "segment_id").
		Values(experiment.UserId, experiment.SegmentId)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "experiment.Create",
		QueryRaw: query,
	}

	if _, err := r.client.PG().Exec(ctx, q, args...); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, experiment *experimentModel.Experiment) (int64, error) {
	builder := sq.Delete(experimentTableName).
		PlaceholderFormat(sq.Dollar).
		Where(
			sq.And{
				sq.Eq{
					"user_id": experiment.UserId,
				},
				sq.Eq{
					"segment_id": experiment.SegmentId,
				},
			},
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := pg.Query{
		Name:     "experiment.Delete",
		QueryRaw: query,
	}

	ct, err := r.client.PG().Exec(ctx, q, args...)
	if err != nil {
		return 0, err
	}

	return ct.RowsAffected(), nil
}
