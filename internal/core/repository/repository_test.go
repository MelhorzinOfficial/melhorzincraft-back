package repository

import (
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/repository/mock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

type StructT struct {
	Id   int `gorm:"primaryKey"`
	Name string
}

func TestNewRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	type testCase[T any] struct {
		args    args
		wantErr bool
	}
	tests := map[string]testCase[StructT]{
		"valid": {
			args: args{
				db: func() *gorm.DB {
					db, _ := mock.NewMockDB()
					return db
				}(),
			},
			wantErr: false,
		},
		"invalid": {
			args: args{
				db: nil,
			},
			wantErr: true,
		},
	}
	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			repo, err := NewRepository[StructT](tt.args.db)
			if tt.wantErr {
				assert.Error(t, err, "expected an error but got none")
				assert.Nil(t, repo, "repository should be nil when error occurs")
			} else {
				assert.NoError(t, err, "expected no error but got one")
				assert.NotNil(t, repo, "repository should not be nil when no error")
			}
		})
	}
}
