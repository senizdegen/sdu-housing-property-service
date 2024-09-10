package minio

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewClient(host, port, username, password string, ssl bool) (*minio.Client, error) {
	endpoint := fmt.Sprintf("%s:%s", host, port)
	options := &minio.Options{
		Creds:  credentials.NewStaticV4(username, password, ""),
		Secure: ssl,
	}

	minioClient, err := minio.New(endpoint, options)
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
