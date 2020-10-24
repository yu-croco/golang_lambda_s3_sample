package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
	"yu-croco/go_lambda_s3/app/adapter"
	"yu-croco/go_lambda_s3/app/infra/localImg"
	"yu-croco/go_lambda_s3/app/infra/s3"
	"yu-croco/go_lambda_s3/app/usecase"
)

func Handler(apiRequest events.APIGatewayProxyRequest) error {
	request, converterErr := adapter.NewFileUploadControllerImpl().Exec(apiRequest.Body)
	if converterErr != nil {
		return converterErr
	}

	fileRepository := localImg.NewLocalFileRepositoryImpl()
	s3Repository := s3.NewS3RepositoryImpl()
	jstCurrentUnixTime := int(time.Now().UTC().In(time.FixedZone("Asia/Tokyo", 9*60*60)).Unix())

	if useCaseErr := usecase.
		NewUploadFileToS3UseCase(request, fileRepository, s3Repository).
		Exec(jstCurrentUnixTime); useCaseErr != nil {
		return useCaseErr
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
