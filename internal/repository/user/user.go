package user

import (
	"context"

	"github.com/zd4r/dynamic-user-segmentation/internal/client/pg"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"

	sq "github.com/Masterminds/squirrel"
)

const (
	userTableName = `"user"`
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

func (r *Repository) Delete(ctx context.Context, user *userModel.User) (int64, error) {
	builder := sq.Delete(userTableName).
		PlaceholderFormat(sq.Dollar).
		Where(
			sq.Eq{
				"id": user.Id,
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
