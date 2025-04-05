package config

import (
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/infra/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	dimagedelivery "github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/delivery"
	dimageuc "github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/usecase"
)

func httpHandler() fx.Option {
	return fx.Provide(
		func(srv *http.Server) *gin.RouterGroup {
			return srv.Group("/api/v1")
		},
	)
}

func registerRoutes() fx.Option {
	return fx.Invoke(
		registerDockerRoutes,
	)
}

func registerDockerRoutes(e *gin.RouterGroup, dockerImageUC *dimageuc.ImageUC) {
	group := e.Group("/docker")

	dimagedelivery.RegisterRoutes(group, dockerImageUC)
}
