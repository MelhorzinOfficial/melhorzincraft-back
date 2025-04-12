package config

import (
	dimageuc "github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/usecase"
	gtemplateuc "github.com/MelhorzinOfficial/melhorzincraft-back/internal/gtemplate/usecase"
	"go.uber.org/fx"
)

func usecases() fx.Option {
	return fx.Provide(
		dimageuc.New,
		gtemplateuc.New,
	)
}
