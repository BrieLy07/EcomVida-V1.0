package s3

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var Cliente *s3.S3
var Bucket string

func InitS3() {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String(os.Getenv("S3_ENDPOINT")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("S3_ACCESS_KEY"),
			os.Getenv("S3_SECRET_KEY"),
			"",
		),
		S3ForcePathStyle: aws.Bool(true), // necesario para MinIO
	})
	if err != nil {
		log.Fatal("Error creando sesión S3:", err)
	}
	Cliente = s3.New(sess)
	Bucket = os.Getenv("S3_BUCKET")
}
