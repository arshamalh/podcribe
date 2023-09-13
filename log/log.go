package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

func init() {
	globalLogger = zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(
				zap.NewDevelopmentEncoderConfig(),
			), zapcore.AddSync(os.Stdout),
			zap.DebugLevel,
		),
	)
}

func GetGlobalLogger() *zap.Logger {
	return globalLogger
}

func Info(msg string, fields ...zap.Field) {
	globalLogger.Info(msg, fields...)
}
