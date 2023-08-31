package segment

import (
	"context"

	"github.com/zd4r/dynamic-user-segmentation/internal/client/pg"
	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"

	sq "github.com/Masterminds/squirrel"
)

const (
	segmentTableName = `segment`
)

type Repository struct {
	client pg.Client
}

func NewRepository(client pg.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) Create(ctx context.Context, segment *segmentModel.Segment) error {
	builder := sq.Insert(segmentTableName).
		PlaceholderFormat(sq.Dollar).
		Columns("slug").
		Values(segment.Slug)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "segment.Create",
		QueryRaw: query,
	}

	if _, err := r.client.PG().Exec(ctx, q, args...); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetBySlug(ctx context.Context, segment *segmentModel.Segment) (*segmentModel.Segment, error) {
	builder := sq.Select("id", "slug").
		PlaceholderFormat(sq.Dollar).
		From(segmentTableName).
		Where(
			sq.Eq{
				"slug": segment.Slug,
			},
		).Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := pg.Query{
		Name:     "segment.GetBySlug",
		QueryRaw: query,
	}

	if err := r.client.PG().ScanOne(ctx, segment, q, args...); err != nil {
		return nil, err
	}

	return segment, nil
}

func (r *Repository) Delete(ctx context.Context, segment *segmentModel.Segment) (int64, error) {
	builder := sq.Delete(segmentTableName).
		PlaceholderFormat(sq.Dollar).
		Where(
			sq.Eq{
				"slug": segment.Slug,
			},
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := pg.Query{
		Name:     "segment.Delete",
		QueryRaw: query,
	}

	ct, err := r.client.PG().Exec(ctx, q, args...)
	if err != nil {
		return 0, err
	}

	return ct.RowsAffected(), nil
}
