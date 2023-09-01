package cron

import (
	"context"
	"time"

	"go.uber.org/zap"
)

const (
	cleanInterval = 1 * time.Minute
)

type DBCleaner struct {
	experimentService experimentService
	shutdown          chan struct{}
	log               *zap.Logger
}

func NewDBCleaner(experimentService experimentService, log *zap.Logger) *DBCleaner {
	return &DBCleaner{
		experimentService: experimentService,
		shutdown:          make(chan struct{}, 1),
		log:               log,
	}
}

func (c *DBCleaner) Start() {
	ctx := context.TODO()
	go func() {
		ticker := time.NewTicker(cleanInterval)
		for {
			select {
			case <-ticker.C:
				rowsAffected, err := c.experimentService.DeleteAllExpired(ctx)
				if err != nil {
					c.log.Error("c.experimentService.DeleteAllExpired", zap.Error(err))
				}
				if rowsAffected != 0 {
					c.log.Info("c.experimentService.DeleteAllExpired", zap.Int64("rows_deleted", rowsAffected))
				}
			case <-c.shutdown:
				c.log.Info("shutdown")
				return
			}
		}
	}()
}

func (c *DBCleaner) Shutdown() {
	c.shutdown <- struct{}{}
}
