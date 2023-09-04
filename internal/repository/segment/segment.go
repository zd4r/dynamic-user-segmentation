package segment

import (
	"context"
	"fmt"

	"github.com/zd4r/dynamic-user-segmentation/internal/client/pg"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"

	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"

	sq "github.com/Masterminds/squirrel"
)

const (
	segmentTableName    = `segment`
	userTableName       = `"user"`
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

func (r *Repository) Create(ctx context.Context, slug string) error {
	builder := sq.Insert(segmentTableName).
		PlaceholderFormat(sq.Dollar).
		Columns("slug").
		Values(slug)

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

func (r *Repository) GetBySlug(ctx context.Context, slug string) (*segmentModel.Segment, error) {
	builder := sq.Select("id", "slug").
		PlaceholderFormat(sq.Dollar).
		From(segmentTableName).
		Where(
			sq.Eq{
				"slug": slug,
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

	segment := new(segmentModel.Segment)
	if err := r.client.PG().ScanOne(ctx, segment, q, args...); err != nil {
		return nil, err
	}

	return segment, nil
}

func (r *Repository) GetUsersBySlug(ctx context.Context, slug string) ([]userModel.User, error) {
	builder := sq.Select(fmt.Sprintf("%s.id", userTableName)).
		PlaceholderFormat(sq.Dollar).
		From(userTableName).
		Join(fmt.Sprintf("%s ON %s.id= %s.user_id", experimentTableName, userTableName, experimentTableName)).
		Join(fmt.Sprintf("%s ON %s.segment_id = %s.id", segmentTableName, experimentTableName, segmentTableName)).
		Where(sq.Eq{
			fmt.Sprintf("%s.slug", segmentTableName): slug,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := pg.Query{
		Name:     "segment.GetUsers",
		QueryRaw: query,
	}

	var users []userModel.User
	if err := r.client.PG().ScanAll(ctx, &users, q, args...); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) DeleteBySlug(ctx context.Context, slug string) (int64, error) {
	builder := sq.Delete(segmentTableName).
		PlaceholderFormat(sq.Dollar).
		Where(
			sq.Eq{
				"slug": slug,
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
