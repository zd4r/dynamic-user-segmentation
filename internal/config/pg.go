package config

import (
	"errors"
	"os"
)

const (
	pgDSNEnvName = "PG_DSN"
)

type PGConfig interface {
	DSN() string
}

var _ PGConfig = (*pgConfig)(nil)

type pgConfig struct {
	dsn string
}

func NewPGConfig() (*pgConfig, error) {
	dsn := os.Getenv(pgDSNEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("pg DSN not found")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
