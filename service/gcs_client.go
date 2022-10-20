package service

import (
	"context"
	"downlod-file-gcs/interfaces"
	"fmt"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type GcsClient struct {
	client *storage.Client
}

func NewGcsClient(client *storage.Client) interfaces.GcsClient {
	return &GcsClient{client: client}
}

func (c *GcsClient) ListFilesWithPrefix(ctx context.Context, bucket, prefix string) ([]*storage.ObjectAttrs, error) {

	it := c.client.Bucket(bucket).Objects(ctx, &storage.Query{
		Prefix:    prefix,
		Delimiter: prefix,
	})
	objects := make([]*storage.ObjectAttrs, 0)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error getting files from %v/%v: %v", bucket, prefix, err)
		}
		if attrs.Name[len(attrs.Name)-1] != '/' {
			objects = append(objects, attrs)
		}
	}

	return objects, nil
}

func (c *GcsClient) DownloadFile(ctx context.Context, bucket, object string) ([]byte, error) {

	rc, err := c.client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("download error %v: %v", object, err)
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("download error %v: %v", object, err)
	}

	return data, nil
}

func (c *GcsClient) UploadFile(ctx context.Context, bucket, object string, content []byte) error {
	return nil
}
