package repository

import (
	"errors"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/dimage"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/repository"
	"gorm.io/gorm"
)

func New(db *gorm.DB) (dimage.Repository, error) {
	if err := db.AutoMigrate(&dimage.DockerImage{}); err != nil {
		return nil, err
	}

	rr, err := repository.NewRepository[dimage.DockerImage](db)
	if err != nil {
		return nil, err
	}

	return &repo{
		Repository: rr,
	}, nil
}

type repo struct {
	repository.Repository[dimage.DockerImage]
}

func (r *repo) ExistByTagAndRepository(tag, repository string) (bool, error) {
	var count int64
	err := r.DB().Model(&dimage.DockerImage{}).
		Where("tag = ? AND repository = ?", tag, repository).
		Count(&count).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return count > 0, nil
}
