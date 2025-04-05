package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"log"
)

func main() {
	cli, err := client.NewClientWithOpts(
		client.WithHost("unix:///Users/peliciari/.docker/run/docker.sock"), // Explicitamente define o socket
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		log.Fatalf("Erro ao criar cliente Docker: %v", err)
	}

	config := &container.Config{
		Image: "nginx",
	}

	resp, err := cli.ContainerCreate(context.Background(), config, nil, nil, nil, "meu_nginx")
	if err != nil {
		log.Fatalf("Erro ao criar contêiner: %v", err)
	}

	if err := cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{}); err != nil {
		log.Fatalf("Erro ao iniciar contêiner: %v", err)
	}

	fmt.Printf("Contêiner %s criado e iniciado com sucesso!\n", resp.ID[:12])
}
