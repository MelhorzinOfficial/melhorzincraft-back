package http

import (
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/infra/http/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Config struct {
	Port int `mapstructure:"port"`
}

func NewServer(cfg *Config) *gin.Engine {
	engine := gin.Default()

	engine.Use(middleware.CorsMiddleware())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return engine
}
