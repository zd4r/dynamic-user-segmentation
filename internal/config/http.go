package config

import (
	"errors"
	"os"
)

const (
	httpPortEnvName = "HTTP_PORT"
)

type HTTPConfig interface {
	Port() string
}

var _ HTTPConfig = (*httpConfig)(nil)

type httpConfig struct {
	port string
}

func NewHTTPConfig() (*httpConfig, error) {
	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("http host not found")
	}

	return &httpConfig{
		port: port,
	}, nil
}

func (cfg *httpConfig) Port() string {
	return cfg.port
}
