package minio

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Config struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
}

var client *minio.Client

type Point struct {
}

func (Point) Upload(file io.Reader, toBucket string, withURL string) error {
	return Upload(file, toBucket, withURL)
}

func (Point) Down(urlFile string, fromBucket string) (to io.Reader, err error) {
	return Down(urlFile, fromBucket)
}

func InitPointWith(conf Config) (p Point, err error) {

	useSSL := false

	client, err = minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKeyID, conf.SecretAccessKey, ""),
		Secure: useSSL,
	})

	return
}

// InitWith Initialize minio client object.
func InitWith(conf Config) (err error) {

	useSSL := false

	client, err = minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKeyID, conf.SecretAccessKey, ""),
		Secure: useSSL,
	})

	return
}

func Upload(file io.Reader, toBucket string, withURL string) error {

	_, err := client.PutObject(context.Background(), toBucket, withURL, file, -1, minio.PutObjectOptions{})

	return err
}

func Down(urlFile string, fromBucket string) (to io.Reader, err error) {

	return client.GetObject(context.Background(), fromBucket, urlFile, minio.GetObjectOptions{})
}
