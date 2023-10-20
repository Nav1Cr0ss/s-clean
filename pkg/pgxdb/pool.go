package pgxdb

import (
	"context"
	libLogger "github.com/Nav1Cr0ss/s-user-storage/pkg/logger/zerolog"
	"github.com/jackc/pgx/v5/tracelog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pkgName = "postgres"
	log     = libLogger.NewDefaultComponentLogger(pkgName)
)

func ParseConfig(c PoolConfig) *pgxpool.Config {
	logger := libLogger.NewDefaultComponentLogger(pkgName)

	pgConfig, err := pgxpool.ParseConfig(c.GetConnString())

	if err != nil {
		logger.Fatal().Err(err).Msg("failed to parse config")
	}

	pgxLogLevel, err := tracelog.LogLevelFromString(c.GetLogLevel())

	if err != nil {
		pgxLogLevel = tracelog.LogLevelNone
	}

	pgConfig.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   NewLogger(logger),
		LogLevel: pgxLogLevel,
	}

	return pgConfig
}

// NewPool is constructor for pgxpool.Pool
func NewPool(c PoolConfig) *pgxpool.Pool {

	pgxPoolConfig := ParseConfig(c)

	var pg *pgxpool.Pool
	var err error

	i := 0
	ticker := time.NewTicker(time.Second)

	for ; ; <-ticker.C {
		i++
		pg, err = pgxpool.NewWithConfig(context.Background(), pgxPoolConfig)
		if err == nil || i > 60 {
			break
		}
	}

	if err != nil {
		log.Fatal().Err(err).Msg("Unable to create connection pool")
	}

	ticker.Stop()

	return pg
}
