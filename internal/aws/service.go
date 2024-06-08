package aws

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct {
	Client *S3Client
}

var S3_URL = "https://%s.s3.amazonaws.com/%s"

func (s *S3Service) UploadFile(file *multipart.File, filename string) (*string, error) {
	_, err := s.Client.Client.PutObjectWithContext(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(s.Client.Bucket),
		Key:    &filename,
		Body:   *file,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	url := fmt.Sprintf(S3_URL, s.Client.Bucket, filename)
	return &url, nil
}

func InitS3Service() *S3Service {
	return &S3Service{}
}

func (s S3Service) WithClient(client *S3Client) S3Service {
	s.Client = client
	return s
}
