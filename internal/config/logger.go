package config

import (
	"errors"
	"os"
)

const (
	logLevelEnvName = "LOG_LEVEL"
)

type LoggerConfig interface {
	LogLevel() string
}

var _ LoggerConfig = (*loggerConfig)(nil)

type loggerConfig struct {
	logLevel string
}

func NewLoggerLevel() (*loggerConfig, error) {
	logLevel := os.Getenv(logLevelEnvName)
	if len(logLevel) == 0 {
		return nil, errors.New("logger level not found")
	}

	return &loggerConfig{
		logLevel: logLevel,
	}, nil
}

func (cfg *loggerConfig) LogLevel() string {
	return cfg.logLevel
}
