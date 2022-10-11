package service

import (
	"context"
	"downlod-file-gcs/interfaces"
	"fmt"

	"cloud.google.com/go/storage"
)

type GcsClient struct {
	client *storage.Client
}

func NewGcsClient(client *storage.Client) interfaces.GcsClient {
	return &GcsClient{client: client}
}

func (c *GcsClient) DownloadFile(ctx context.Context, bucket, object string) ([]byte, error) {
	return []byte{1}, fmt.Errorf("download failed")
}
