package repository

import (
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/dimage"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/repository"
	"gorm.io/gorm"
)

func New(db *gorm.DB) dimage.Repository {
	return &repo{
		Repository: repository.NewRepository[dimage.DockerImage](db),
		db:         db,
	}
}

type repo struct {
	repository.Repository[dimage.DockerImage]

	db *gorm.DB
}
