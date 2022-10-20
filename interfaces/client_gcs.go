package interfaces

import (
	"context"

	"cloud.google.com/go/storage"
)

type GcsClient interface {
	ListFilesWithPrefix(ctx context.Context, bucket, prefix string) ([]*storage.ObjectAttrs, error)
	DownloadFile(ctx context.Context, bucket, object string) ([]byte, error)
	UploadFile(ctx context.Context, bucket, object string, content []byte) error
}
