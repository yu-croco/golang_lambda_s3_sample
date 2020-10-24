package usecase

import (
	"yu-croco/go_lambda_s3/app/domain"
	"yu-croco/go_lambda_s3/app/domain/model"
)

type uploadFileToS3UseCaseImpl struct {
	Request             *model.Request
	LocalFileRepository domain.LocalFileRepository
	S3Repository        domain.S3Repository
}

type UploadFileToS3UseCase interface {
	Exec(jstCurrentUnixTime int) error
}

func NewUploadFileToS3UseCase(request *model.Request,
	fileRepository domain.LocalFileRepository,
	s3Repository domain.S3Repository) UploadFileToS3UseCase {

	return uploadFileToS3UseCaseImpl{
		Request:             request,
		LocalFileRepository: fileRepository,
		S3Repository:        s3Repository,
	}
}

func (impl uploadFileToS3UseCaseImpl) Exec(jstCurrentUnixTime int) error {
	if localFileErr := impl.LocalFileRepository.Add(impl.Request, jstCurrentUnixTime); localFileErr != nil {
		return localFileErr
	}

	if s3Err := impl.S3Repository.Add(impl.Request, jstCurrentUnixTime); s3Err != nil {
		return s3Err
	}

	return nil
}
