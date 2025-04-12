package gtemplate

import (
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/dimage"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/repository"
	"time"
)

//go:generate mockgen -destination=./mock/mock_gtemplate.go -package=mock github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/gtemplate Repository

type GameTemplate struct {
	DockerImage *dimage.DockerImage

	Id            int    `gorm:"primaryKey"`
	Name          string `gorm:"column:name"`
	DockerCompose []byte `gorm:"column:docker_compose"`

	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"index"`
}

func (g *GameTemplate) TableName() string {
	return "game_templates"
}

type Repository interface {
	repository.Repository[GameTemplate]
}
