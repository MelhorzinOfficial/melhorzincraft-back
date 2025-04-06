package docker

import (
	"context"
	"fmt"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/docker"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	dclient "github.com/docker/docker/client"
	"io"
	"os"
)

type Config struct {
	Host string `mapstructure:"host"`
}

func NewClient(cfg *Config) (docker.Docker, error) {
	cli, err := dclient.NewClientWithOpts(
		dclient.WithHost(cfg.Host),
		dclient.WithAPIVersionNegotiation(),
	)

	if err != nil {
		return nil, fmt.Errorf("error creating docker client: %w", err)
	}

	return &client{cli: cli}, nil
}

type client struct {
	cli *dclient.Client
}

// TODO: implement PullOtions
func (c *client) ImagePull(ctx context.Context, tag string) error {
	out, err := c.cli.ImagePull(ctx, tag, image.PullOptions{})
	if err != nil {
		return fmt.Errorf("error pulling image: %w", err)
	}
	defer out.Close()

	io.Copy(os.Stdout, out)

	return nil
}

// TODO: implement RemoveOptions
func (c *client) ImageDelete(ctx context.Context, tag string) error {
	_, err := c.cli.ImageRemove(ctx, tag, image.RemoveOptions{})
	if err != nil {
		return fmt.Errorf("error deleting image: %w", err)
	}

	return nil
}

func (c *client) RunContainer(ctx context.Context, name string, image string) (string, error) {
	resp, err := c.cli.ContainerCreate(ctx, &container.Config{
		Image: image,
	}, nil, nil, nil, name)
	if err != nil {
		return "", err
	}

	if err := c.cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return "", err
	}

	return resp.ID, nil
}
