package gcp

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func (gcp *GcpClient) ListBuckets(ctx context.Context, projectId string) (buckets []string, err error) {
	it := gcp.client.Buckets(ctx, projectId)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		buckets = append(buckets, battrs.Name)
	}
	return buckets, nil
}

func (gcp *GcpClient) ListObjects(ctx context.Context, bucket string, path string) (objs []string, err error) {
	storageQuery := &storage.Query{Prefix: path, Delimiter: "/"}

	it := gcp.client.Bucket(bucket).Objects(ctx, storageQuery)
	for {
		oattrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		objs = append(objs, oattrs.Name)
	}
	return objs, nil
}
