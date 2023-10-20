package postgres

import (
	logger "github.com/Nav1Cr0ss/s-user-storage/pkg/logger/zerolog"
	"github.com/Nav1Cr0ss/s-user-storage/pkg/pgxdb"
)

type Repository struct {
	db  *pgxdb.Database
	log logger.Logger
}

// NewRepository opens new db pool connection
func NewRepository(db *pgxdb.Database) *Repository {
	return &Repository{
		db:  db,
		log: logger.NewDefaultComponentLogger("postgres"),
	}
}
