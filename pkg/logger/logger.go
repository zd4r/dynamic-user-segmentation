package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(logLevel zapcore.Level, nameSpace string) *zap.Logger {
	conf := zap.NewDevelopmentEncoderConfig()
	conf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(conf),
		zapcore.AddSync(os.Stdout),
		logLevel,
	))
	return logger.Named(nameSpace)
}
