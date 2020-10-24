package util

import (
	"strconv"
	"yu-croco/go_lambda_s3/app/domain/model"
)

func LambdaLocalFilePath(jstCurrentUnixTime int) string {
	return "/tmp/" + strconv.Itoa(jstCurrentUnixTime) + ".jpg"
}

func S3FilePath(req *model.Request, jstCurrentUnixTime int) string {
	return "/images/" + basePath(req, jstCurrentUnixTime)
}

func basePath(req *model.Request, jstCurrentUnixTime int) string {
	return strconv.Itoa(req.UserId) + "/" + strconv.Itoa(jstCurrentUnixTime) + ".jpg"
}
