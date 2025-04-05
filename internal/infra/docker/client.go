package docker

import (
	"fmt"
	"github.com/docker/docker/client"
)

// TODO: Implement a Docker client that can be used to interact with the Docker API.
func NewClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(
		client.WithHost("unix:///var/run/docker.sock"),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating docker client: %w", err)
	}
	return cli, nil
}
