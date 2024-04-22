package gcp

import (
	"context"
	"time"

	"cloud.google.com/go/storage"
	"github.com/achintya-7/ez-storage/model"
	"google.golang.org/api/iterator"
)

// returns the list of buckets in a project
func (gcp *GcpClient) ListBuckets(ctx context.Context, projectId string) (buckets []string, err error) {
	it := gcp.client.Buckets(ctx, projectId)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, &model.GcpError{Err: err}
		}
		buckets = append(buckets, battrs.Name)
	}
	return buckets, nil
}

// returns the list of objects in a bucket at a given path
func (gcp *GcpClient) ListObjects(ctx context.Context, bucket string, path string) (objs []string, err error) {
	storageQuery := &storage.Query{Prefix: path, Delimiter: "/"}

	it := gcp.client.Bucket(bucket).Objects(ctx, storageQuery)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, &model.GcpError{Err: err}
		}
		objs = append(objs, attrs.Name)
	}
	return objs, nil
}

// returns the size of a folder in MB in a bucket at a given path
func (gcp *GcpClient) GetPathSize(ctx context.Context, bucket string, path string) (size int64, err error) {
	storageQuery := &storage.Query{Prefix: path, Delimiter: "/"}

	it := gcp.client.Bucket(bucket).Objects(ctx, storageQuery)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, &model.GcpError{Err: err}
		}

		mbSize := attrs.Size / 1024 / 1024

		size += mbSize
	}
	return size, nil
}

// deletes a folder in a bucket at a given path
func (gcp *GcpClient) DeleteFolder(ctx context.Context, bucket, path string) (err error) {
	err = gcp.client.Bucket(bucket).Object(path).Delete(ctx)
	if err != nil {
		return &model.GcpError{Err: err}
	}

	return nil
}

// returns a signed download URL for a file in a bucket at a given path
func (gcp *GcpClient) GetSignedDownloadURL(ctx context.Context, bucket, path string, expiry time.Time) (url string, err error) {
	opts := &storage.SignedURLOptions{
		Scheme: storage.SigningSchemeV4,
		Method: "GET",
		Headers: []string{
			"Content-Type:application/octet-stream",
		},
		Expires: expiry,
	}

	url, err = storage.SignedURL(bucket, path, opts)
	if err != nil {
		return "", &model.GcpError{Err: err}
	}

	return url, nil
}

func (gcp *GcpClient) GetSignedUploadUrl(ctx context.Context, bucket, path string, expiry time.Time) (url string, err error) {
	opts := &storage.SignedURLOptions{
		Scheme: storage.SigningSchemeV4,
		Method: "PUT",
		Headers: []string{
			"Content-Type:application/octet-stream",
		},
		Expires: expiry,
	}

	url, err = storage.SignedURL(bucket, path, opts)
	if err != nil {
		return "", &model.GcpError{Err: err}
	}

	return url, nil
}

func (gcp *GcpClient) GetAttributes(ctx context.Context, bucket, path string) (attrs *model.ObjAttrs, err error) {
	attr, err := gcp.client.Bucket(bucket).Object(path).Attrs(ctx)
	if err != nil {
		return nil, &model.GcpError{Err: err}
	}

	attrs = &model.ObjAttrs{
		Name:    attr.Name,
		Size:    attr.Size,
		Created: attr.Created,
		Updatd:  attr.Updated,
	}

	return attrs, nil
}
