package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/repository"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/usecase"
	"log"
	"net/http"

	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/delivery"

	_ "github.com/MelhorzinOfficial/melhorzincraft-back/docs"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/infra/db"
	infra_http "github.com/MelhorzinOfficial/melhorzincraft-back/internal/infra/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// @title MelhorzinCraft API
// @version 1.0
// @description Docker image management API for MelhorzinCraft
// @host localhost:8080
// @BasePath /api/v1
func main() {
	app := fx.New(
		fx.Provide(
			newConfig,
			func(cfg *AppConfig) *infra_http.Config {
				return &cfg.HTTP
			},
			func(cfg *AppConfig) *db.Config {
				return &cfg.DB
			},
			infra_http.NewServer,
			db.NewClient,
			newHTTPServer,
			usecase.New,
			repository.New,
		),
		fx.Invoke(
			delivery.RegisterRoutes,
			startHTTPServer,
			func(db *gorm.DB) {
				log.Println("Database connected")
			},
		),
	)

	app.Run()
}

// TODO: refactor this
type AppConfig struct {
	HTTP infra_http.Config
	DB   db.Config
}

func newConfig() *AppConfig {
	return &AppConfig{
		HTTP: infra_http.Config{
			Port: 8080,
		},
		DB: db.Config{
			Url: "sqlite://melhorzincraft.db",
		},
	}
}

func startHTTPServer(lifecycle fx.Lifecycle, server *http.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("Error starting server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}

func newHTTPServer(cfg *AppConfig, engine *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HTTP.Port),
		Handler: engine,
	}
}
