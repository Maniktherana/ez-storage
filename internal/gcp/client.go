package gcp

import (
	"context"
	"sync"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type GcpClient struct {
	client *storage.Client
	mu     sync.Mutex
}

func NewGcpClient(ctx context.Context, opts ...option.ClientOption) (*GcpClient, error) {
	client, err := storage.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return &GcpClient{client: client}, nil
}
