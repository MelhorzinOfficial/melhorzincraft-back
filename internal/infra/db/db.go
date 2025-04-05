package db

import (
	"errors"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Url string `mapstructure:"url"`
}

func NewClient(cfg *Config) (*gorm.DB, error) {
	if cfg == nil || cfg.Url == "" {
		return nil, errors.New("db config is nil")
	}

	url := cfg.Url

	dialector, err := getDialect(url)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDialect(url string) (gorm.Dialector, error) {
	switch {
	case strings.HasPrefix(url, "postgres://"), strings.HasPrefix(url, "postgresql://"):
		return postgres.Open(url), nil

	case strings.HasPrefix(url, "mysql://"):
		dsn := strings.TrimPrefix(url, "mysql://")
		return mysql.Open(dsn), nil

	case strings.HasPrefix(url, "sqlite://"):
		cleanUrl := strings.TrimPrefix(url, "sqlite://")
		return sqlite.Open(cleanUrl), nil

	case strings.HasPrefix(url, "file:"):
		return sqlite.Open(url), nil

	case !strings.Contains(url, "://") && !strings.Contains(url, ":"):
		return sqlite.Open(url), nil
	}

	return nil, errors.New("dialect don't supported")
}
