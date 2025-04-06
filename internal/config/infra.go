package config

import (
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/infra/db"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/infra/docker"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/infra/http"
	"go.uber.org/fx"
)

func infra() fx.Option {
	return fx.Provide(
		db.NewClient,
		http.NewServer,
		docker.NewClient,
	)
}
