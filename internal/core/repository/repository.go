package repository

import (
	"context"
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Save(ctx context.Context, obj *T) error
	FindById(ctx context.Context, id int) (*T, error)
	Delete(ctx context.Context, obj *T) error
}

func NewRepository[T any](db *gorm.DB) Repository[T] {
	return &repository[T]{
		db: db,
	}
}

type repository[T any] struct {
	db *gorm.DB
}

func (r *repository[T]) Save(ctx context.Context, obj *T) error {
	return r.db.WithContext(ctx).Save(obj).Error
}

func (r *repository[T]) FindById(ctx context.Context, id int) (*T, error) {
	var obj T
	err := r.db.WithContext(ctx).First(&obj, id).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func (r *repository[T]) Delete(ctx context.Context, obj *T) error {
	return r.db.WithContext(ctx).Delete(obj).Error
}
