package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zd4r/dynamic-user-segmentation/internal/client/pg"
	"github.com/zd4r/dynamic-user-segmentation/internal/config"
	"github.com/zd4r/dynamic-user-segmentation/pkg/closer"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	httpConfig config.HTTPConfig

	pgClient pg.Client
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GetPGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GetHTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) GetPGClient(ctx context.Context) pg.Client {
	if s.pgClient == nil {
		pgCfg, err := pgxpool.ParseConfig(s.GetPGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to parse pg config from dsn: %s", err.Error())
		}

		pgClient, err := pg.NewClient(ctx, pgCfg)
		if err != nil {
			log.Fatalf("failet to get pg client: %s", err.Error())
		}

		if err := pgClient.PG().Ping(ctx); err != nil {
			log.Fatalf("failed to ping db: %s", err.Error())
		}

		closer.Add(pgClient.Close)

		s.pgClient = pgClient
	}

	return s.pgClient
}
