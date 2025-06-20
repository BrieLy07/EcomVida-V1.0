package utils

import (
	"context"
	"io"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var clienteMinio *minio.Client
var bucket string

func init() {
	var err error
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	bucket = os.Getenv("MINIO_BUCKET")

	clienteMinio, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("Error al conectar con MinIO: %v", err)
	}

	ctx := context.Background()
	found, err := clienteMinio.BucketExists(ctx, bucket)
	if err != nil {
		log.Fatalf("Error verificando bucket: %v", err)
	}
	if !found {
		err = clienteMinio.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("No se pudo crear el bucket: %v", err)
		}
	}
}

func SubirAminio(file io.Reader, nombre string, contentType string) error {
	_, err := clienteMinio.PutObject(context.Background(), bucket, nombre, file, -1, minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

func GenerarURLImagen(nombre string) (string, error) {
	reqParams := url.Values{}
	url, err := clienteMinio.PresignedGetObject(context.Background(), bucket, nombre, time.Minute*10, reqParams)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}
