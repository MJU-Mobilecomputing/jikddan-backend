package aws

import (
	"errors"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3client *S3Client
	once     sync.Once
)

func InitS3Client(bucket string, key string, privateKey string) (*S3Client, error) {
	if s3client == nil {
		once.Do(func() {
			_ = setS3Client(bucket, key, privateKey)
		})
		return s3client, nil
	} else {
		return nil, errors.New("s3client already configured")
	}
}

func setS3Client(bucket string, key string, privateKey string) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-2"),
		Credentials: credentials.NewStaticCredentials(key, privateKey, ""),
	}))
	client := s3.New(sess)
	s3client = &S3Client{
		Client:  client,
		Session: sess,
		Bucket:  bucket,
	}
	return nil
}
