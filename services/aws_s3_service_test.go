package services

import (
	"os"
	"path/filepath"
	"testing"
)

var awsS3Service = AWSS3Service{}

func TestAWSS3Service_UploadFileUploadFile(t *testing.T) {

	var err error
	awsS3Service, err = awsS3Service.Initialize()

	if nil != err {
		t.Errorf("Failed to initialize aws s3 service: %s", err.Error())
	}

	pwd, _ := os.Getwd()

	filePath := pwd + "/" + os.Getenv("TEST_IMAGE_FILE")

	fileData, err := os.Open(filePath)
	if nil != err {
		t.Errorf("Failed to read file with path %s : %s", os.Getenv("TEST_IMAGE_FILE"), err.Error())
	}

	_, file := filepath.Split(filePath)

	result, err := awsS3Service.UploadFile(fileData, os.Getenv("S3_DEFAULT_PATH"), file)

	if "" == *result.ETag {
		t.Errorf("Failed to upload file to path %s : %s", os.Getenv("S3_DEFAULT_PATH")+"/"+file, err.Error())
	}

}

func TestAWSS3Service_DeleteFile(t *testing.T) {

}

func TestAWSS3Service_GetFile(t *testing.T) {

}
