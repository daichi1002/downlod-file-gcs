package interfaces

import "context"

type GcsClient interface {
	DownloadFile(ctx context.Context, bucket, object string) ([]byte, error)
}
