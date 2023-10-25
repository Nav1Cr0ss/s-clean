package start

import (
	"context"
	"github.com/Nav1Cr0ss/s-user-storage/internal/adapters"
	"github.com/Nav1Cr0ss/s-user-storage/internal/adapters/repository/postgres"
	"github.com/Nav1Cr0ss/s-user-storage/internal/app"
	"github.com/Nav1Cr0ss/s-user-storage/internal/config"
	"github.com/Nav1Cr0ss/s-user-storage/internal/ports"
	"github.com/Nav1Cr0ss/s-user-storage/internal/service"
	"github.com/Nav1Cr0ss/s-user-storage/pkg/configurator"
	logger "github.com/Nav1Cr0ss/s-user-storage/pkg/logger/zerolog"
	fxzerolog "github.com/Nav1Cr0ss/s-user-storage/pkg/logger/zerolog/adapters/fx"
	"github.com/Nav1Cr0ss/s-user-storage/pkg/pgxdb"
	"github.com/Nav1Cr0ss/s-user-storage/pkg/xecho"
	"go.uber.org/fx"
)

func Server() {
	fx.New(
		fx.WithLogger(
			fxzerolog.Init(),
		),
		fx.Provide(
			fx.Annotate(
				configurator.NewConfig[config.Config],
				fx.As(new(xecho.HTTPServerConfig)),
				fx.As(new(pgxdb.PoolConfig)),
				fx.As(new(app.AppConfig)),
				fx.As(new(logger.LoggerConfig)),
			),
			logger.NewDefaultLogger,
			pgxdb.New,
			fx.Annotate(postgres.NewRepository, fx.As(new(adapters.Repository))),
			app.New,
		),
		ports.NewServer(),
		fx.Invoke(
			func(lc fx.Lifecycle, log logger.Logger, repo adapters.Repository, userApp *service.Service) {
				lc.Append(
					fx.Hook{
						OnStart: func(ctx context.Context) error {
							log.Info().Msg("start app")
							return nil
						},
						OnStop: func(ctx context.Context) error {
							log.Info().Msg("shutdown app")
							return nil
						},
					},
				)
			},
		),
	).Run()
}
