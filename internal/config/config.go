package config

import (
	"context"

	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/infra/db"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/infra/http"
	"go.uber.org/fx"
)

type Config struct {
	fx.Out

	DB   *db.Config   `mapstructure:"db"`
	Http *http.Config `mapstructure:"http"`
}

func (cfg *Config) Get() Config {
	return *cfg
}

func (cfg *Config) Start() {
	app := fx.New(
		fx.Provide(cfg.Get),

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
