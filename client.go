package ezstorage

import (
	"context"
	"errors"
	"time"

	"github.com/achintya-7/ez-storage/internal/gcp"
	"github.com/achintya-7/ez-storage/model"
)

type StorageConfig struct {
	Type    string
	Context context.Context
}

type StorageFunctioner interface {
	ListBuckets(ctx context.Context, projectId string) (buckets []string, err error)
	ListObjects(ctx context.Context, bucket, path string) (objs []string, err error)
	GetPathSize(ctx context.Context, bucket string, path string) (size int64, err error)
	DeleteFolder(ctx context.Context, bucket, path string) (err error)
	GetSignedDownloadURL(ctx context.Context, bucket, path string, expiry time.Time) (url string, err error)
	GetSignedUploadUrl(ctx context.Context, bucket, path string, expiry time.Time) (url string, err error)
	GetAttributes(ctx context.Context, bucket, path string) (attrs *model.ObjAttrs, err error)
}

// todo: Implement AWS client
// todo: Add option to pass opts to client
func NewClient(config StorageConfig) (storageClient StorageFunctioner, err error) {
	if config.Context == nil {
		config.Context = context.Background()
	}

	switch config.Type {
	case model.GCP:
		gcpClient, err := gcp.NewGcpClient(config.Context)
		if err != nil {
			return nil, err
		}

		return gcpClient, nil

	case model.AWS:

	default:
		return nil, errors.New("invalid storage type in config")
	}

	return nil, errors.New("invalid storage type in config")
}
