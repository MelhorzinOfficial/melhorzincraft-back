package config

import (
	"go.uber.org/fx"

	dimagerepo "github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/repository"
)

func repositories() fx.Option {
	return fx.Provide(
		dimagerepo.New,
	)
}
