package usecase

import (
	"context"
	"errors"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/dimage"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/dimage/mock"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/response"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/usecase/dto"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/oserror"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestImageUC_Create(t *testing.T) {
	type fields struct {
		repo dimage.Repository
		v    *validator.Validate
	}
	type args struct {
		req *dto.CreateRequest
	}

	ctrl := gomock.NewController(t)
	m := mock.NewMockRepository(ctrl)

	tests := map[string]struct {
		fields fields
		args   args
		setup  func(m *mock.MockRepository)
		want   *response.Response[dto.CreateResponse]
	}{
		"create success": {
			fields: fields{
				repo: m,
				v:    validator.New(),
			},
			args: args{
				req: &dto.CreateRequest{
					Tag:        "latest",
					Repository: "test/test",
				},
			},
			setup: func(m *mock.MockRepository) {
				m.EXPECT().ExistByTagAndRepository(gomock.Any(), gomock.Any()).Return(false, nil)
				m.EXPECT().Create(gomock.Any(), &dimage.DockerImage{
					Tag:        "latest",
					Repository: "test/test",
				}).DoAndReturn(func(ctx context.Context, obj *dimage.DockerImage) error {
					obj.Id = 1
					obj.CreatedAt = time.Unix(0, 0).UTC()
					obj.UpdatedAt = time.Unix(0, 0).UTC()
					return nil
				})
			},
			want: &response.Response[dto.CreateResponse]{
				Code:    200,
				Error:   false,
				Message: "success",
				Data: dto.CreateResponse{
					DockerImage: &dto.DockerImage{
						Id:         1,
						Tag:        "latest",
						Repository: "test/test",
						CreatedAt:  time.Unix(0, 0).UTC(),
						UpdatedAt:  time.Unix(0, 0).UTC(),
					},
				},
			},
		},
		"create conflict": {
			fields: fields{
				repo: m,
				v:    validator.New(),
			},
			args: args{
				req: &dto.CreateRequest{
					Tag:        "latest",
					Repository: "test/test",
				},
			},
			setup: func(m *mock.MockRepository) {
				m.EXPECT().ExistByTagAndRepository(gomock.Any(), gomock.Any()).Return(true, nil)
			},
			want: &response.Response[dto.CreateResponse]{
				Code:    409,
				Error:   true,
				Message: oserror.ErrDockerImageAlreadyExists.Error(),
			},
		},
		"server error": {
			fields: fields{
				repo: m,
				v:    validator.New(),
			},
			args: args{
				req: &dto.CreateRequest{
					Tag:        "latest",
					Repository: "test/test",
				},
			},
			setup: func(m *mock.MockRepository) {
				m.EXPECT().ExistByTagAndRepository(gomock.Any(), gomock.Any()).Return(false, errors.New("server error"))
			},
			want: &response.Response[dto.CreateResponse]{
				Code:    500,
				Error:   true,
				Message: oserror.ErrServer.Error(),
			},
		},
	}
	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			i := &ImageUC{
				repo: tt.fields.repo,
				v:    tt.fields.v,
			}

			tt.setup(m)

			got := i.Create(t.Context(), tt.args.req)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestImageUC_Update(t *testing.T) {
	type fields struct {
		repo dimage.Repository
		v    *validator.Validate
	}
	type args struct {
		req *dto.UpdateRequest
	}

	ctrl := gomock.NewController(t)
	m := mock.NewMockRepository(ctrl)

	tests := map[string]struct {
		fields fields
		args   args
		setup  func(m *mock.MockRepository)
		want   *response.Response[dto.UpdateResponse]
	}{
		"create success": {
			fields: fields{
				repo: m,
				v:    validator.New(),
			},
			args: args{
				req: &dto.UpdateRequest{
					Id:         1,
					Tag:        "latest",
					Repository: "test/test",
				},
			},
			setup: func(m *mock.MockRepository) {
				di := &dimage.DockerImage{
					Id:         1,
					Tag:        "latest",
					Repository: "test/test",
					CreatedAt:  time.Unix(0, 0).UTC(),
					UpdatedAt:  time.Unix(0, 0).UTC(),
					DeletedAt:  nil,
				}

				m.EXPECT().FindById(gomock.Any(), 1).Return(di, nil)

				m.EXPECT().Update(gomock.Any(), di).
					DoAndReturn(func(ctx context.Context, obj *dimage.DockerImage) error {
						obj.Tag = "updated"
						obj.Repository = "updated/test"
						return nil
					})
			},
			want: &response.Response[dto.UpdateResponse]{
				Code:    200,
				Error:   false,
				Message: "success",
				Data: dto.UpdateResponse{
					DockerImage: &dto.DockerImage{
						Id:         1,
						Tag:        "updated",
						Repository: "updated/test",
						CreatedAt:  time.Unix(0, 0).UTC(),
						UpdatedAt:  time.Unix(0, 0).UTC(),
					},
				},
			},
		},
		"not found": {
			fields: fields{
				repo: m,
				v:    validator.New(),
			},
			args: args{
				req: &dto.UpdateRequest{
					Id:         1,
					Tag:        "latest",
					Repository: "test/test",
				},
			},
			setup: func(m *mock.MockRepository) {
				m.EXPECT().FindById(gomock.Any(), 1).Return(nil, oserror.ErrNotFound)
			},
			want: &response.Response[dto.UpdateResponse]{
				Code:    404,
				Error:   true,
				Message: oserror.ErrNotFound.Error(),
			},
		},
		"server error": {
			fields: fields{
				repo: m,
				v:    validator.New(),
			},
			args: args{
				req: &dto.UpdateRequest{
					Id:         1,
					Tag:        "latest",
					Repository: "test/test",
				},
			},
			setup: func(m *mock.MockRepository) {
				m.EXPECT().FindById(gomock.Any(), 1).Return(nil, errors.New("server error"))
			},
			want: &response.Response[dto.UpdateResponse]{
				Code:    500,
				Error:   true,
				Message: oserror.ErrServer.Error(),
			},
		},
	}
	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			i := &ImageUC{
				repo: tt.fields.repo,
				v:    tt.fields.v,
			}

			tt.setup(m)

			got := i.Update(t.Context(), tt.args.req)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestImageUC_Show(t *testing.T) {
	type fields struct {
		repo dimage.Repository
		v    *validator.Validate
	}
	type args struct {
		req *dto.ShowRequest
	}

	ctrl := gomock.NewController(t)
	m := mock.NewMockRepository(ctrl)

	tests := map[string]struct {
		fields fields
		args   args
		setup  func(m *mock.MockRepository)
		want   *response.Response[dto.ShowResponse]
	}{
		"create success": {
			fields: fields{
				repo: m,
				v:    validator.New(),
			},
			args: args{
				req: &dto.ShowRequest{
					Id: 1,
				},
			},
			setup: func(m *mock.MockRepository) {
				m.EXPECT().FindById(gomock.Any(), 1).Return(&dimage.DockerImage{
					Id:         1,
					Tag:        "latest",
					Repository: "test/test",
					CreatedAt:  time.Unix(0, 0).UTC(),
					UpdatedAt:  time.Unix(0, 0).UTC(),
					DeletedAt:  nil,
				}, nil)
			},
			want: &response.Response[dto.ShowResponse]{
				Code:    200,
				Error:   false,
				Message: "success",
				Data: dto.ShowResponse{
					DockerImage: &dto.DockerImage{
						Id:         1,
						Tag:        "latest",
						Repository: "test/test",
						CreatedAt:  time.Unix(0, 0).UTC(),
						UpdatedAt:  time.Unix(0, 0).UTC(),
					},
				},
			},
		},
		"not found": {
			fields: fields{
				repo: m,
				v:    validator.New(),
			},
			args: args{
				req: &dto.ShowRequest{
					Id: 1,
				},
			},
			setup: func(m *mock.MockRepository) {
				m.EXPECT().FindById(gomock.Any(), 1).Return(nil, oserror.ErrNotFound)
			},
			want: &response.Response[dto.ShowResponse]{
				Code:    404,
				Error:   true,
				Message: oserror.ErrNotFound.Error(),
			},
		},
		"server error": {
			fields: fields{
				repo: m,
				v:    validator.New(),
			},
			args: args{
				req: &dto.ShowRequest{
					Id: 1,
				},
			},
			setup: func(m *mock.MockRepository) {
				m.EXPECT().FindById(gomock.Any(), 1).Return(nil, oserror.ErrServer)
			},
			want: &response.Response[dto.ShowResponse]{
				Code:    500,
				Error:   true,
				Message: oserror.ErrServer.Error(),
			},
		},
	}
	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			i := &ImageUC{
				repo: tt.fields.repo,
				v:    tt.fields.v,
			}

			tt.setup(m)

			got := i.Show(t.Context(), tt.args.req)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestImageUC_Delete(t *testing.T) {
	type fields struct {
		repo dimage.Repository
		v    *validator.Validate
	}
	type args struct {
		req *dto.DeleteRequest
	}

	ctrl := gomock.NewController(t)
	m := mock.NewMockRepository(ctrl)

	tests := map[string]struct {
		fields fields
		args   args
		setup  func(m *mock.MockRepository)
		want   *response.Response[dto.DeleteResponse]
	}{
		"create success": {
			fields: fields{
				repo: m,
				v:    validator.New(),
			},
			args: args{
				req: &dto.DeleteRequest{
					Id: 1,
				},
			},
			setup: func(m *mock.MockRepository) {
				m.EXPECT().FindById(gomock.Any(), 1).Return(&dimage.DockerImage{
					Id:         1,
					Tag:        "latest",
					Repository: "test/test",
					CreatedAt:  time.Unix(0, 0).UTC(),
					UpdatedAt:  time.Unix(0, 0).UTC(),
					DeletedAt:  nil,
				}, nil)

				m.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			want: &response.Response[dto.DeleteResponse]{
				Code:    200,
				Error:   false,
				Message: "success",
				Data:    dto.DeleteResponse{},
			},
		},
		"not found": {
			fields: fields{
				repo: m,
				v:    validator.New(),
			},
			args: args{
				req: &dto.DeleteRequest{
					Id: 1,
				},
			},
			setup: func(m *mock.MockRepository) {
				m.EXPECT().FindById(gomock.Any(), 1).Return(nil, oserror.ErrNotFound)

			},
			want: &response.Response[dto.DeleteResponse]{
				Code:    404,
				Error:   true,
				Message: oserror.ErrNotFound.Error(),
			},
		},
		"server error": {
			fields: fields{
				repo: m,
				v:    validator.New(),
			},
			args: args{
				req: &dto.DeleteRequest{
					Id: 1,
				},
			},
			setup: func(m *mock.MockRepository) {
				m.EXPECT().FindById(gomock.Any(), 1).Return(nil, oserror.ErrServer)
			},
			want: &response.Response[dto.DeleteResponse]{
				Code:    500,
				Error:   true,
				Message: oserror.ErrServer.Error(),
			},
		},
	}
	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			i := &ImageUC{
				repo: tt.fields.repo,
				v:    tt.fields.v,
			}

			tt.setup(m)

			got := i.Delete(t.Context(), tt.args.req)
			assert.Equal(t, tt.want, got)
		})
	}
}
