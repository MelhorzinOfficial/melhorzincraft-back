package config

import (
	"context"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/infra/db"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/infra/docker"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/infra/http"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/logger"
	"go.uber.org/fx"
)

type Config struct {
	fx.Out

	DB     *db.Config     `mapstructure:"db"`
	Http   *http.Config   `mapstructure:"http"`
	Docker *docker.Config `mapstructure:"docker"`
	Logger *logger.Config `mapstructure:"logger"`
}

func (cfg *Config) Get() Config {
	return *cfg
}

func (cfg *Config) Start() {
	app := fx.New(
		fx.Provide(cfg.Get, logger.New),

		repositories(),
		usecases(),
		infra(),
		httpHandler(),

		registerRoutes(),
		httpServer(),
	)

	app.Run()
}

func httpServer() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle, srv *http.Server) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go srv.Start()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		})
	})
}
