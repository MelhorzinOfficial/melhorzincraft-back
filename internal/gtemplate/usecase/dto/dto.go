package dto

import "time"

type GameTemplate struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	DockerCompose string    `json:"docker_compose"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateRequest struct {
	Name          string `json:"name"`
	DockerCompose string `json:"docker_compose"`
}

type CreateResponse struct {
	*GameTemplate
}

type UpdateRequest struct {
	Id            int    `uri:"id" swaggerignore:"true"`
	Name          string `json:"name"`
	DockerCompose string `json:"docker_compose"`
}

type UpdateResponse struct {
	*GameTemplate
}

type ShowRequest struct {
	Id int `uri:"id"`
}

type ShowResponse struct {
	*GameTemplate
}

type DeleteRequest struct {
	Id int `uri:"id"`
}

type DeleteResponse struct {
}
