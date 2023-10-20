package pgxdb

import (
	logger "github.com/Nav1Cr0ss/s-user-storage/pkg/logger/zerolog"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PoolConfig interface {
	GetConnString() string
	GetLogLevel() string
}

type Database struct {
	*pgxpool.Pool
	log logger.Logger
}

func New(c PoolConfig) *Database {
	d := &Database{
		Pool: NewPool(c),
		log:  logger.NewDefaultComponentLogger("db"),
	}
	return d
}
