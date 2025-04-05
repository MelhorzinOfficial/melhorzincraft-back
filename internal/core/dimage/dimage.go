package dimage

import (
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/repository"
	"time"
)

type DockerImage struct {
	Id         int        `gorm:"primaryKey"`
	Tag        string     `gorm:"column:tag"`
	Repository string     `gorm:"column:repository"`
	CreatedAt  time.Time  `gorm:"autoCreateTime, column:created_at"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime, column:updated_at"`
	DeletedAt  *time.Time `gorm:"index"`
}

func (i *DockerImage) TableName() string {
	return "docker_images"
}

type Repository interface {
	repository.Repository[DockerImage]

	ExistByTagAndRepository(tag, repository string) (bool, error)
}
