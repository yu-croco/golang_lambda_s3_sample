package localImg

import (
	"encoding/base64"
	"os"
	"yu-croco/go_lambda_s3/app/domain"
	"yu-croco/go_lambda_s3/app/domain/model"
	"yu-croco/go_lambda_s3/app/util"
)

type localFileRepositoryImpl struct {
	Request *model.Request
}

func NewLocalFileRepositoryImpl() domain.LocalFileRepository {
	return &localFileRepositoryImpl{}
}

func (impl *localFileRepositoryImpl) Add(req *model.Request, jstCurrentUnixTime int) error {
	data, decodeErr := base64.StdEncoding.DecodeString(req.Image)
	if decodeErr != nil {
		return decodeErr
	}

	file, err := os.Create(util.LambdaLocalFilePath(jstCurrentUnixTime))
	if err != nil {
		return err
	}

	defer file.Close()
	file.Write(data)

	return nil
}
