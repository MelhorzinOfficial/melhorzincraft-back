package oserror

import "errors"

var (
	ErrInvalidDatabase          = errors.New("invalid database")
	ErrNotFound                 = errors.New("not found")
	ErrAlreadyExists            = errors.New("already exists")
	ErrDockerImageAlreadyExists = errors.New("docker image already exists")
	ErrServer                   = errors.New("server error")
)
