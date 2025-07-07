package s3

import (
	"bytes"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func SubirAObjeto(nombre string, datos []byte) (string, error) {
	_, err := Cliente.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(Bucket),
		Key:    aws.String(nombre),
		Body:   bytes.NewReader(datos),
	})
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("%s/%s/%s", Cliente.Endpoint, Bucket, nombre)
	return url, nil
}

func ObtenerObjeto(nombre string) ([]byte, error) {
	salida, err := Cliente.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(Bucket),
		Key:    aws.String(nombre),
	})
	if err != nil {
		return nil, err
	}
	defer salida.Body.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, salida.Body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func EliminarObjeto(nombre string) error {
	_, err := Cliente.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(Bucket),
		Key:    aws.String(nombre),
	})
	return err
}
