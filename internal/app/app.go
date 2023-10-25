package app

import (
	"github.com/Nav1Cr0ss/s-user-storage/internal/adapters"
	"github.com/Nav1Cr0ss/s-user-storage/internal/app/user"
	"github.com/Nav1Cr0ss/s-user-storage/internal/service"
	logger "github.com/Nav1Cr0ss/s-user-storage/pkg/logger/zerolog"
)

type AppConfig interface {
	GetRSAPublicKey() string
}

type app struct {
	repo adapters.Repository
	c    AppConfig
	log  logger.Logger
}

func New(
	c AppConfig,
	repo adapters.Repository,
) *service.Service {

	var log = logger.NewDefaultComponentLogger("app")

	_ = log
	//pubKey, err := os.ReadFile(cfg.RSAPublicKey)
	//if err != nil {
	//	log.Panic().Err(err).Msg("failed to read public key")
	//}
	//
	//encrypt.SetPublicKey(pubKey)

	return &service.Service{
		User: user.New(repo, c),
	}
}
