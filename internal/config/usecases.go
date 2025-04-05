package config

import (
	dimageuc "github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/usecase"
	"go.uber.org/fx"
)

func usecases() fx.Option {
	return fx.Provide(
		dimageuc.New,
	)
}
