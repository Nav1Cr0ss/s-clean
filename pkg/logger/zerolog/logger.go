package logger

import (
	"context"
	"github.com/Nav1Cr0ss/s-user-storage/internal/config"
	"github.com/Nav1Cr0ss/s-user-storage/pkg/configurator"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LoggerConfig interface {
	GetLogConsole() bool
	GetLogLevel() string
}

// Logger is wrapper struct around logger.Logger that adds some custom functionality
type Logger struct {
	zerolog.Logger
}

// Printf is implementation of fx.Printer
func (l Logger) Printf(s string, args ...interface{}) {
	l.Info().Msgf(s, args...)
}

// NewLogger return logger instance
func NewLogger(component string, output io.Writer, c LoggerConfig) Logger {

	level, err := zerolog.ParseLevel(c.GetLogLevel())
	if err != nil {
		level = zerolog.InfoLevel
	}

	if c.GetLogConsole() {
		output = zerolog.NewConsoleWriter()
	} else if output == nil {
		output = os.Stdout
	}

	l := zerolog.
		New(output).
		Level(level).
		//Hook(SeverityHook{}).
		//Hook(ReportLocationHook{skip: cfg.LogConsole}).
		//Hook(TypeHook{skip: cfg.LogConsole}).
		With().Timestamp()

	if level == zerolog.DebugLevel || level == zerolog.TraceLevel {
		l = l.Caller()
	}

	if component != "" {
		l = l.Str("component", component)
	}

	return Logger{l.Logger()}
}

// NewDefaultLogger return default logger instance
func NewDefaultLogger() Logger {
	return NewLogger("", os.Stdout, configurator.NewConfig[config.Config]())
}

// NewDefaultComponentLogger return default logger instance with custom component
func NewDefaultComponentLogger(component string) Logger {
	return NewLogger(component, os.Stdout, configurator.NewConfig[config.Config]())
}

// NewDefaultConsoleLogger return default logger instance
func NewDefaultConsoleLogger() Logger {
	return NewLogger("", zerolog.NewConsoleWriter(), configurator.NewConfig[config.Config]())
}

func Ctx(ctx context.Context) *zerolog.Logger {
	l := log.Ctx(ctx)

	if l.GetLevel() == zerolog.Disabled {
		*l = NewDefaultConsoleLogger().Logger
	}

	return l
}
