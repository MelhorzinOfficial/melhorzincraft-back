package docker

import "context"

type Docker interface {
	ImagePull(ctx context.Context, tag string) error
	ImageDelete(ctx context.Context, tag string) error
	RunContainer(ctx context.Context, name string, image string) (string, error)
}
