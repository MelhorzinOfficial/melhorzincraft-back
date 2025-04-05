package repository

import (
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/dimage"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/repository"
	"gorm.io/gorm"
)

func New(db *gorm.DB) (dimage.Repository, error) {
	rr, err := repository.NewRepository[dimage.DockerImage](db)
	if err != nil {
		return nil, err
	}

	return &repo{
		Repository: rr,
		db:         db,
	}, nil
}

type repo struct {
	repository.Repository[dimage.DockerImage]

	db *gorm.DB
}
