package usecase

import (
	"context"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/dimage"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/response"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/usecase/dto"
	"github.com/go-playground/validator/v10"
)

func New(repo dimage.Repository) *ImageUC {
	return &ImageUC{
		repo: repo,
	}
}

type ImageUC struct {
	repo dimage.Repository

	v *validator.Validate
}

func (i *ImageUC) Save(ctx context.Context, req *dto.CreateRequest) response.Response[dto.CreateResponse] {
	if err := i.v.Struct(req); err != nil {
		return response.NewBadRequest[dto.CreateResponse](err.Error())
	}

	image := dimage.DockerImage{
		Tag:        req.Tag,
		Repository: req.Repository,
	}

	if err := i.repo.Save(ctx, &image); err != nil {
		return response.NewBadRequest[dto.CreateResponse](err.Error())
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

func (i *ImageUC) Show(ctx context.Context, req *dto.ShowRequest) response.Response[dto.ShowResponse] {
	if err := i.v.Struct(req); err != nil {
		return response.NewBadRequest[dto.ShowResponse](err.Error())
	}

	image, err := i.repo.FindById(ctx, req.Id)
	if err != nil {
		return response.NewBadRequest[dto.ShowResponse](err.Error())
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

func (i *ImageUC) Delete(ctx context.Context, req *dto.DeleteRequest) response.Response[dto.DeleteResponse] {
	if err := i.v.Struct(req); err != nil {
		return response.NewBadRequest[dto.DeleteResponse](err.Error())
	}

	image, err := i.repo.FindById(ctx, req.Id)
	if err != nil {
		return response.NewNotFound[dto.DeleteResponse](err.Error())
	}

	if err := i.repo.Delete(ctx, image); err != nil {
		return response.NewInternalServerError[dto.DeleteResponse](err.Error())
	}

	return response.NewSuccess[dto.DeleteResponse](dto.DeleteResponse{
		Message: "Docker image deleted successfully",
	})
}
