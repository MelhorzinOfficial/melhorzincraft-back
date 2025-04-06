package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	_ "github.com/MelhorzinOfficial/melhorzincraft-back/docs"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/config"
	"github.com/spf13/viper"
)

// @title MelhorzinCraft API
// @version 1.0
// @description Docker image management API for MelhorzinCraft
// @host localhost:8080
// @BasePath /api/v1
func main() {
	var cfg config.Config

	v := initViper()

	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error loading configurations: %v", err)
	}

	cfg.Start()
}

func initViper() *viper.Viper {
	v := viper.New()

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./configs")

	setDefaults(v)

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			fmt.Println("Configuration file not found, using environment variables and default values")
		}
	} else {
		fmt.Printf("Using configuration file: %s\n", v.ConfigFileUsed())
	}

	return v
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("http.address", ":8080")

	v.SetDefault("db.url", "sqlite://:memory:")

	v.SetDefault("docker.host", "unix:///Users/peliciari/.docker/run/docker.sock") // TODO: check this
}
