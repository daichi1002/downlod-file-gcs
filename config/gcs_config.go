package config

import (
	"context"
	"downlod-file-gcs/interfaces"
	"downlod-file-gcs/service"
	"fmt"

	"cloud.google.com/go/storage"
)

func ConnectGCS(ctx context.Context) (interfaces.GcsClient, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	return service.NewGcsClient(client), err
}
