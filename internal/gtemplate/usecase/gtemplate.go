package usecase

import (
	"context"
	"errors"

	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/gtemplate"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/logger"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/response"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/gtemplate/usecase/dto"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/oserror"
	"github.com/go-playground/validator/v10"
)

func New(repo gtemplate.Repository, l logger.Logger) *TemplateUC {
	return &TemplateUC{
		repo: repo,
		v:    validator.New(),
		l:    l,
	}
}

type TemplateUC struct {
	repo gtemplate.Repository

	v *validator.Validate
	l logger.Logger
}

func (t *TemplateUC) Create(ctx context.Context, req *dto.CreateRequest) *response.Response[dto.CreateResponse] {
	t.l.Debug("create template usecase", "request", req)

	if err := t.v.Struct(req); err != nil {
		t.l.Error("validation error", "error", err.Error())
		return response.NewBadRequest[dto.CreateResponse](err.Error())
	}

	template := gtemplate.GameTemplate{
		Name:          req.Name,
		DockerCompose: []byte(req.DockerCompose),
	}

	if err := t.repo.Create(ctx, &template); err != nil {
		t.l.Error("template repository error", "error", err.Error())
		return response.NewInternalServerError[dto.CreateResponse]()
	}

	return response.NewSuccess[dto.CreateResponse](dto.CreateResponse{
		GameTemplate: &dto.GameTemplate{
			Id:            template.Id,
			Name:          template.Name,
			DockerCompose: string(template.DockerCompose),
			CreatedAt:     template.CreatedAt.UTC(),
			UpdatedAt:     template.UpdatedAt.UTC(),
		},
	})
}

func (t *TemplateUC) Update(ctx context.Context, req *dto.UpdateRequest) *response.Response[dto.UpdateResponse] {
	t.l.Debug("update template usecase", "request", req)

	if err := t.v.Struct(req); err != nil {
		t.l.Error("validation error", "error", err.Error())
		return response.NewBadRequest[dto.UpdateResponse](err.Error())
	}

	template, err := t.repo.FindById(ctx, req.Id)
	if err != nil {
		t.l.Error("template repository error", "error", err.Error())
		if errors.Is(err, oserror.ErrNotFound) {
			return response.NewNotFound[dto.UpdateResponse]()
		}
		return response.NewInternalServerError[dto.UpdateResponse]()
	}

	if req.Name != "" && req.Name != template.Name {
		template.Name = req.Name
	}

	if req.DockerCompose != "" && req.DockerCompose != string(template.DockerCompose) {
		template.DockerCompose = []byte(req.DockerCompose)
	}

	if err := t.repo.Update(ctx, template); err != nil {
		t.l.Error("template repository error", "error", err.Error())
		return response.NewInternalServerError[dto.UpdateResponse]()
	}

	return response.NewSuccess[dto.UpdateResponse](dto.UpdateResponse{
		GameTemplate: &dto.GameTemplate{
			Id:            template.Id,
			Name:          template.Name,
			DockerCompose: string(template.DockerCompose),
			CreatedAt:     template.CreatedAt.UTC(),
			UpdatedAt:     template.UpdatedAt.UTC(),
		},
	})
}

func (t *TemplateUC) Show(ctx context.Context, req *dto.ShowRequest) *response.Response[dto.ShowResponse] {
	t.l.Debug("show template usecase", "request", req)

	if err := t.v.Struct(req); err != nil {
		t.l.Error("validation error", "error", err.Error())
		return response.NewBadRequest[dto.ShowResponse](err.Error())
	}

	template, err := t.repo.FindById(ctx, req.Id)
	if err != nil {
		t.l.Error("template repository error", "error", err.Error())

		if errors.Is(err, oserror.ErrNotFound) {
			return response.NewNotFound[dto.ShowResponse]()
		}

		return response.NewInternalServerError[dto.ShowResponse]()
	}

	return response.NewSuccess[dto.ShowResponse](dto.ShowResponse{
		GameTemplate: &dto.GameTemplate{
			Id:            template.Id,
			Name:          template.Name,
			DockerCompose: string(template.DockerCompose),
			CreatedAt:     template.CreatedAt.UTC(),
			UpdatedAt:     template.UpdatedAt.UTC(),
		},
	})
}

func (t *TemplateUC) Delete(ctx context.Context, req *dto.DeleteRequest) *response.Response[dto.DeleteResponse] {
	t.l.Debug("delete template usecase", "request", req)

	if err := t.v.Struct(req); err != nil {
		t.l.Error("validation error", "error", err.Error())
		return response.NewBadRequest[dto.DeleteResponse](err.Error())
	}

	template, err := t.repo.FindById(ctx, req.Id)
	if err != nil {
		t.l.Error("template repository error", "error", err.Error())

		if errors.Is(err, oserror.ErrNotFound) {
			return response.NewNotFound[dto.DeleteResponse]()
		}
		return response.NewInternalServerError[dto.DeleteResponse]()
	}

	if err := t.repo.Delete(ctx, template); err != nil {
		t.l.Error("template repository error", "error", err.Error())
		return response.NewInternalServerError[dto.DeleteResponse]()
	}

	return response.NewSuccess[dto.DeleteResponse](dto.DeleteResponse{})
}
