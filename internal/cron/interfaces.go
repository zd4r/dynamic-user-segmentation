package cron

import (
	"context"
)

type experimentService interface {
	DeleteAllExpired(ctx context.Context) (int64, error)
}
