package report

import (
	"context"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/zd4r/dynamic-user-segmentation/internal/client/pg"
	reportModel "github.com/zd4r/dynamic-user-segmentation/internal/model/report"
)

const (
	reportTableName = `report`
)

type Repository struct {
	client pg.Client
}

func NewRepository(client pg.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (r *Repository) CreateBatchRecord(ctx context.Context, records []reportModel.Record) error {
	builder := sq.Insert(reportTableName).
		PlaceholderFormat(sq.Dollar).
		Columns("user_id", "segment_slug", "action")

	for _, record := range records {
		builder = builder.Values(record.UserId, record.SegmentSlug, record.Action)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "report.CreateRecord",
		QueryRaw: query,
	}

	if _, err := r.client.PG().Exec(ctx, q, args...); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetRecordsInIntervalByUser(ctx context.Context, userId int, from, to time.Time) ([]reportModel.Record, error) {
	builder := sq.Select("user_id", "segment_slug", "action", "date").
		PlaceholderFormat(sq.Dollar).
		From(reportTableName).
		Where(
			sq.And{
				sq.Eq{
					"user_id": userId,
				},
				sq.GtOrEq{
					"date": from,
				},
				sq.LtOrEq{
					"date": to,
				},
			},
		).OrderBy("date DESC")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	log.Println(query)

	q := pg.Query{
		Name:     "report.GetRecordsInInterval",
		QueryRaw: query,
	}

	records := make([]reportModel.Record, 0)
	if err := r.client.PG().ScanAll(ctx, &records, q, args...); err != nil {
		return nil, err
	}

	return records, nil
}
