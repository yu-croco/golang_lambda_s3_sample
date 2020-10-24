package adapter

import (
	"encoding/json"
	"yu-croco/go_lambda_s3/app/domain/model"
)

type requestConverterImpl struct{}

type requestConverter interface {
	Exec(requestBody string) (*model.Request, error)
}

func NewFileUploadControllerImpl() requestConverter {
	return requestConverterImpl{}
}

func (impl requestConverterImpl) Exec(requestBody string) (*model.Request, error) {
	jsonBytes := []byte(requestBody)
	req := new(model.Request)

	if unmarshalErr := json.Unmarshal(jsonBytes, req); unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return req, nil
}
