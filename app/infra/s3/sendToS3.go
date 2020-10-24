package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"yu-croco/go_lambda_s3/app/domain"
	"yu-croco/go_lambda_s3/app/domain/model"
	"yu-croco/go_lambda_s3/app/util"
)

type s3RepositoryImpl struct{}

func NewS3RepositoryImpl() domain.S3Repository {
	return s3RepositoryImpl{}
}

func (impl s3RepositoryImpl) Add(req *model.Request, jstCurrentUnixTime int) error {
	file, openErr := os.Open(util.LambdaLocalFilePath(jstCurrentUnixTime))
	if openErr != nil {
		return openErr
	}

	defer file.Close()
	if uploadErr := impl.upload(req, file, jstCurrentUnixTime); uploadErr != nil {
		return uploadErr
	}

	return nil
}

func (impl s3RepositoryImpl) upload(req *model.Request, file *os.File, jstCurrentUnixTime int) error {
	uploader := s3manager.NewUploader(impl.newSession())

	if _, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(util.S3FilePath(req, jstCurrentUnixTime)),
		Body:   file,
		ACL:    aws.String("public-read"),
	}); err != nil {
		return err
	}

	return nil
}

func (impl s3RepositoryImpl) newSession() *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
}
