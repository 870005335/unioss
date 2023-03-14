package unioss

import "github.com/minio/minio-go"

type Minio struct {
	config *Config
	client *minio.Client
	bucket string
}
