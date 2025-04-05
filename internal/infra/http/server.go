package http

import (
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/infra/http/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type Config struct {
	Address string `mapstructure:"address"`
}

func NewServer(cfg *Config) *Server {
	engine := gin.Default()

	engine.Use(middleware.CorsMiddleware())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return &Server{
		Addr:    cfg.Address,
		Handler: engine,
	}
}

type Server struct {
	Addr    string
	Handler *gin.Engine
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.Addr, s.Handler)
}

func (s *Server) Group(prefix string) *gin.RouterGroup {
	return s.Handler.Group(prefix)
}
