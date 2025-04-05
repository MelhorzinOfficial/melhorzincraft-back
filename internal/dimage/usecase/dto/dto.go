package dto

import "time"

type DockerImage struct {
	Id         int       `json:"id"`
	Tag        string    `json:"tag"`
	Repository string    `json:"repository"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateRequest struct {
	Tag        string `json:"tag"`
	Repository string `json:"repository"`
}

type CreateResponse struct {
	*DockerImage
}

type UpdateRequest struct {
	Id         int    `uri:"id"`
	Tag        string `json:"tag"`
	Repository string `json:"repository"`
}

type UpdateResponse struct {
	*DockerImage
}

type ShowRequest struct {
	Id int `uri:"id"`
}

type ShowResponse struct {
	*DockerImage
}

type DeleteRequest struct {
	Id int `uri:"id"`
}

type DeleteResponse struct {
}
