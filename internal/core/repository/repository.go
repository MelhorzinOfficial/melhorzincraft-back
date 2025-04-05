package repository

import (
	"context"
	"errors"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/oserror"
	"gorm.io/gorm"
)

type Repository[T any] interface {
	DB() *gorm.DB

	Create(ctx context.Context, obj *T) error
	Update(ctx context.Context, obj *T) error
	FindById(ctx context.Context, id int) (*T, error)
	Delete(ctx context.Context, obj *T) error
}

func NewRepository[T any](db *gorm.DB) (Repository[T], error) {
	if db == nil {
		return nil, oserror.ErrInvalidDatabase
	}

	return &repository[T]{
		db: db,
	}, nil
}

type repository[T any] struct {
	db *gorm.DB
}

func (r *repository[T]) DB() *gorm.DB {
	return r.db
}

func (r *repository[T]) Create(ctx context.Context, obj *T) error {
	return r.db.WithContext(ctx).Create(obj).Error
}

func (r *repository[T]) Update(ctx context.Context, obj *T) error {
	return r.db.WithContext(ctx).Updates(obj).Error
}

func (r *repository[T]) FindById(ctx context.Context, id int) (*T, error) {
	var obj T
	err := r.db.WithContext(ctx).First(&obj, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, oserror.ErrNotFound
		}
		return nil, err
	}
	return &obj, nil
}

func (r *repository[T]) Delete(ctx context.Context, obj *T) error {
	if err := r.db.WithContext(ctx).Delete(obj).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return oserror.ErrNotFound
		}
		return err
	}

	return nil
}
