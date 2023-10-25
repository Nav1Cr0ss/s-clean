package pgxdb

import (
	"context"
	libLogger "github.com/Nav1Cr0ss/s-user-storage/pkg/logger/zerolog"

	"github.com/jackc/pgx/v5/tracelog"
)

type Logger struct {
	log libLogger.Logger
}

func NewLogger(log libLogger.Logger) *Logger {
	return &Logger{log: log}
}

func (l *Logger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]interface{}) {
	logger := libLogger.Ctx(ctx).With().Str("component", pkgName).Logger()

	switch level {
	case tracelog.LogLevelTrace:
		logger.Trace().Interface("data", data).Msg(msg)
	case tracelog.LogLevelDebug:
		logger.Debug().Interface("data", data).Msg(msg)
	case tracelog.LogLevelInfo:
		logger.Info().Interface("data", data).Msg(msg)
	case tracelog.LogLevelWarn:
		logger.Warn().Interface("data", data).Msg(msg)
	case tracelog.LogLevelError:
		logger.Error().Interface("data", data).Msg(msg)
	}
}
