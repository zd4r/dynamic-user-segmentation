package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zd4r/dynamic-user-segmentation/internal/client/pg"
	"github.com/zd4r/dynamic-user-segmentation/internal/config"
	segmentRepo "github.com/zd4r/dynamic-user-segmentation/internal/repository/segment"
	userRepo "github.com/zd4r/dynamic-user-segmentation/internal/repository/user"
	segmentSrv "github.com/zd4r/dynamic-user-segmentation/internal/service/segment"
	userSrv "github.com/zd4r/dynamic-user-segmentation/internal/service/user"
	"github.com/zd4r/dynamic-user-segmentation/pkg/closer"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	httpConfig config.HTTPConfig

	userRepository *userRepo.Repository
	userService    *userSrv.Service

	segmentRepository *segmentRepo.Repository
	segmentService    *segmentSrv.Service

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

func (s *serviceProvider) GetUserRepository(ctx context.Context) *userRepo.Repository {
	if s.userRepository == nil {
		s.userRepository = userRepo.NewRepository(s.GetPGClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) GetUserService(ctx context.Context) *userSrv.Service {
	if s.userService == nil {
		s.userService = userSrv.NewService(s.GetUserRepository(ctx))
	}

	return s.userService
}

func (s *serviceProvider) GetSegmentRepository(ctx context.Context) *segmentRepo.Repository {
	if s.segmentRepository == nil {
		s.segmentRepository = segmentRepo.NewRepository(s.GetPGClient(ctx))
	}

	return s.segmentRepository
}

func (s *serviceProvider) GetSegmentService(ctx context.Context) *segmentSrv.Service {
	if s.segmentService == nil {
		s.segmentService = segmentSrv.NewService(s.GetSegmentRepository(ctx))
	}

	return s.segmentService
}
