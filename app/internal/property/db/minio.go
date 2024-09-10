package db

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/senizdegen/sdu-housing/property-service/pkg/logging"
)

type minioDB struct {
	client     *minio.Client
	logger     logging.Logger
	bucketName string
}

type Minio interface {
	Upload(ctx context.Context, imagePath string) (string, error)
}

func NewMinio(client *minio.Client, logger logging.Logger, buckerName string) Minio {
	return &minioDB{
		client:     client,
		logger:     logger,
		bucketName: buckerName,
	}
}

func (m *minioDB) Upload(ctx context.Context, imagePath string) (string, error) {
	return "", nil
}
