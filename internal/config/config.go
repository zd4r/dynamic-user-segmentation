package config

import (
	"context"

	"github.com/joho/godotenv"
)

const dotEnvFileName = ".env"

func Init(_ context.Context) error {
	return godotenv.Load(dotEnvFileName)
}
