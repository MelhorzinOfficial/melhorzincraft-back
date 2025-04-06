package usecase

import (
	"context"
	"errors"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/dimage"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/docker"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/logger"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/response"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/usecase/dto"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/oserror"
	"github.com/go-playground/validator/v10"
)

func New(repo dimage.Repository, l logger.Logger, docker docker.Docker) *ImageUC {
	return &ImageUC{
		repo: repo,
		v:    validator.New(),
		l:    l,
		dc:   docker,
	}
}

type ImageUC struct {
	repo dimage.Repository

	v *validator.Validate
	l logger.Logger

	dc docker.Docker
}

func (i *ImageUC) Create(ctx context.Context, req *dto.CreateRequest) *response.Response[dto.CreateResponse] {
	i.l.Debug("create image usecase", "request", req)

	if err := i.v.Struct(req); err != nil {
		i.l.Error("validation error", "error", err.Error())
		return response.NewBadRequest[dto.CreateResponse](err.Error())
	}

	image := dimage.DockerImage{
		Tag:        req.Tag,
		Repository: req.Repository,
	}

	exists, err := i.repo.ExistByTagAndRepository(image.Tag, image.Repository)
	if err != nil {
		i.l.Error("image repository error", "error", err.Error())
		return response.NewInternalServerError[dto.CreateResponse]()
	}
	i.l.Debug("image repository exists", "exists", exists)

	if exists {
		return response.NewConflict[dto.CreateResponse](oserror.ErrDockerImageAlreadyExists.Error())
	}

	if err := i.dc.ImagePull(ctx, image.GetFullTag()); err != nil {
		i.l.Error("image pull error", "error", err.Error())
		return response.NewInternalServerError[dto.CreateResponse]()
	}

	if err := i.repo.Create(ctx, &image); err != nil {
		i.l.Error("image repository error", "error", err.Error())
		return response.NewInternalServerError[dto.CreateResponse]()
	}

	return response.NewSuccess[dto.CreateResponse](dto.CreateResponse{
		DockerImage: &dto.DockerImage{
			Id:         image.Id,
			Tag:        image.Tag,
			Repository: image.Repository,
			CreatedAt:  image.CreatedAt.UTC(),
			UpdatedAt:  image.UpdatedAt.UTC(),
		},
	})
}

func (i *ImageUC) Update(ctx context.Context, req *dto.UpdateRequest) *response.Response[dto.UpdateResponse] {
	i.l.Debug("update image usecase", "request", req)

	if err := i.v.Struct(req); err != nil {
		i.l.Error("validation error", "error", err.Error())
		return response.NewBadRequest[dto.UpdateResponse](err.Error())
	}

	image, err := i.repo.FindById(ctx, req.Id)
	if err != nil {
		i.l.Error("image repository error", "error", err.Error())
		if errors.Is(err, oserror.ErrNotFound) {
			return response.NewNotFound[dto.UpdateResponse]()
		}
		return response.NewInternalServerError[dto.UpdateResponse]()
	}

	if req.Tag != "" && req.Tag != image.Tag {
		image.Tag = req.Tag
	}

	if req.Repository != "" && req.Repository != image.Repository {
		image.Repository = req.Repository
	}

	if err := i.dc.ImagePull(ctx, image.GetFullTag()); err != nil {
		i.l.Error("image pull error", "error", err.Error())
		return response.NewInternalServerError[dto.UpdateResponse]()
	}

	if err := i.repo.Update(ctx, image); err != nil {
		i.l.Error("image repository error", "error", err.Error())
		return response.NewInternalServerError[dto.UpdateResponse]()
	}

	return response.NewSuccess[dto.UpdateResponse](dto.UpdateResponse{
		DockerImage: &dto.DockerImage{
			Id:         image.Id,
			Tag:        image.Tag,
			Repository: image.Repository,
			CreatedAt:  image.CreatedAt.UTC(),
			UpdatedAt:  image.UpdatedAt.UTC(),
		},
	})
}

func (i *ImageUC) Show(ctx context.Context, req *dto.ShowRequest) *response.Response[dto.ShowResponse] {
	i.l.Debug("show image usecase", "request", req)

	if err := i.v.Struct(req); err != nil {
		i.l.Error("validation error", "error", err.Error())
		return response.NewBadRequest[dto.ShowResponse](err.Error())
	}

	image, err := i.repo.FindById(ctx, req.Id)
	if err != nil {
		i.l.Error("image repository error", "error", err.Error())

		if errors.Is(err, oserror.ErrNotFound) {
			return response.NewNotFound[dto.ShowResponse]()
		}

		return response.NewInternalServerError[dto.ShowResponse]()
	}

	return response.NewSuccess[dto.ShowResponse](dto.ShowResponse{
		DockerImage: &dto.DockerImage{
			Id:         image.Id,
			Tag:        image.Tag,
			Repository: image.Repository,
			CreatedAt:  image.CreatedAt.UTC(),
			UpdatedAt:  image.UpdatedAt.UTC(),
		},
	})
}

func (i *ImageUC) Delete(ctx context.Context, req *dto.DeleteRequest) *response.Response[dto.DeleteResponse] {
	i.l.Debug("delete image usecase", "request", req)

	if err := i.v.Struct(req); err != nil {
		i.l.Error("validation error", "error", err.Error())
		return response.NewBadRequest[dto.DeleteResponse](err.Error())
	}

	image, err := i.repo.FindById(ctx, req.Id)
	if err != nil {
		i.l.Error("image repository error", "error", err.Error())

		if errors.Is(err, oserror.ErrNotFound) {
			return response.NewNotFound[dto.DeleteResponse]()
		}
		return response.NewInternalServerError[dto.DeleteResponse]()
	}

	if err := i.dc.ImageDelete(ctx, image.GetFullTag()); err != nil {
		i.l.Error("image delete error", "error", err.Error())
		return response.NewInternalServerError[dto.DeleteResponse]()
	}

	if err := i.repo.Delete(ctx, image); err != nil {
		i.l.Error("image repository error", "error", err.Error())
		return response.NewInternalServerError[dto.DeleteResponse]()
	}

	return response.NewSuccess[dto.DeleteResponse](dto.DeleteResponse{})
}
