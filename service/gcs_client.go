package service

import (
	"context"
	"downlod-file-gcs/interfaces"
	"fmt"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type GcsClient struct {
	client *storage.Client
}

func NewGcsClient(client *storage.Client) interfaces.GcsClient {
	return &GcsClient{client: client}
}

func (c *GcsClient) ListFilesWithPrefix(ctx context.Context, bucket string) ([]*storage.ObjectAttrs, error) {

	it := c.client.Bucket(bucket).Objects(ctx, &storage.Query{})
	objects := make([]*storage.ObjectAttrs, 0)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("%vからファイル取得エラー: %v", bucket, err)
		}
		if attrs.Name[len(attrs.Name)-1] != '/' {
			objects = append(objects, attrs)
		}
	}

	return objects, nil
}
