package user

import (
	"context"
	"fmt"
	"time"

	"github.com/zd4r/dynamic-user-segmentation/internal/client/pg"
	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"

	sq "github.com/Masterminds/squirrel"
)

const (
	userTableName       = `"user"`
	segmentTableName    = `segment`
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

func (r *Repository) Create(ctx context.Context, user *userModel.User) error {
	builder := sq.Insert(userTableName).
		PlaceholderFormat(sq.Dollar).
		Columns("id").
		Values(user.Id)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "user.Create",
		QueryRaw: query,
	}

	if _, err := r.client.PG().Exec(ctx, q, args...); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetAll(ctx context.Context) ([]userModel.User, error) {
	builder := sq.Select("*").
		From(userTableName)
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := pg.Query{
		Name:     "user.GetAll",
		QueryRaw: query,
	}

	var users []userModel.User
	if err := r.client.PG().ScanAll(ctx, &users, q, args...); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) Delete(ctx context.Context, userId int) (int64, error) {
	builder := sq.Delete(userTableName).
		PlaceholderFormat(sq.Dollar).
		Where(
			sq.Eq{
				"id": userId,
			},
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := pg.Query{
		Name:     "user.Delete",
		QueryRaw: query,
	}

	ct, err := r.client.PG().Exec(ctx, q, args...)
	if err != nil {
		return 0, err
	}

	return ct.RowsAffected(), nil
}

func (r *Repository) GetSegments(ctx context.Context, userId int) ([]segmentModel.Segment, error) {
	builder := sq.Select(fmt.Sprintf("%s.id", segmentTableName), fmt.Sprintf("%s.slug", segmentTableName)).
		PlaceholderFormat(sq.Dollar).
		From(segmentTableName).
		Join(fmt.Sprintf("%s ON %s.id = %s.segment_id", experimentTableName, segmentTableName, experimentTableName)).
		Join(fmt.Sprintf("%s ON %s.id = %s.user_id", userTableName, userTableName, experimentTableName)).
		Where(
			sq.And{
				sq.Eq{
					fmt.Sprintf("%s.id", userTableName): userId,
				},
				sq.Or{
					sq.GtOrEq{
						fmt.Sprintf("%s.expire_at", experimentTableName): time.Now(),
					},
					sq.Eq{
						fmt.Sprintf("%s.expire_at", experimentTableName): nil,
					},
				},
			},
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := pg.Query{
		Name:     "user.GetUserSegments",
		QueryRaw: query,
	}

	var segments []segmentModel.Segment
	err = r.client.PG().ScanAll(ctx, &segments, q, args...)
	if err != nil {
		return nil, err
	}

	return segments, nil
}
