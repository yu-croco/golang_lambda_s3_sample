package domain

import "yu-croco/go_lambda_s3/app/domain/model"

type LocalFileRepository interface {
	Add(req *model.Request, jstCurrentUnixTime int) error
}

type S3Repository interface {
	Add(req *model.Request, jstCurrentUnixTime int) error
}

