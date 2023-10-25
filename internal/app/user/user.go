package user

import (
	"github.com/Nav1Cr0ss/s-user-storage/internal/adapters"
)

type AppConfig interface {
	GetRSAPublicKey() string
}

type UserApp struct {
	repo adapters.Repository
	cfg  AppConfig
}

func New(
	repo adapters.Repository, c AppConfig,
) *UserApp {
	return &UserApp{
		repo: repo,
		cfg:  c,
	}
}
