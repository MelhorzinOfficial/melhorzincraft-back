package repository

import (
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/gtemplate"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/repository"
	"gorm.io/gorm"
)

func New(db *gorm.DB) (gtemplate.Repository, error) {
	if err := db.AutoMigrate(&gtemplate.GameTemplate{}); err != nil {
		return nil, err
	}

	rr, err := repository.NewRepository[gtemplate.GameTemplate](db)
	if err != nil {
		return nil, err
	}

	return &repo{
		Repository: rr,
	}, nil
}

type repo struct {
	repository.Repository[gtemplate.GameTemplate]
}
