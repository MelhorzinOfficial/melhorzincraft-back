package config

import (
	"go.uber.org/fx"

	dimagerepo "github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/repository"
	gtemplate "github.com/MelhorzinOfficial/melhorzincraft-back/internal/gtemplate/repository"
)

func repositories() fx.Option {
	return fx.Provide(
		dimagerepo.New,
		gtemplate.New,
	)
}
