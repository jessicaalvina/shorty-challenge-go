package services

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"testing"
)

var awsS3Service = AWSS3Service{}

func TestAWSS3Service_UploadFileUploadFile(t *testing.T) {

	pwd, _ := os.Getwd()
	godotenv.Load(pwd + "/../.env")

	var err error
	awsS3Service, err = awsS3Service.Initialize()

	if nil != err {
		t.Errorf("Failed to initialize aws s3 service: %s", err.Error())
	}

	filePath := pwd + "/../" + os.Getenv("TEST_IMAGE_FILE")

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

func TestAWSS3Service_GetFile(t *testing.T) {

	pwd, _ := os.Getwd()
	godotenv.Load(pwd + "/../.env")

	var err error
	awsS3Service, err = awsS3Service.Initialize()

	if nil != err {
		t.Errorf("Failed to initialize aws s3 service: %s", err.Error())
	}

	filePath := pwd + "/../" + os.Getenv("TEST_IMAGE_FILE")
	_, file := filepath.Split(filePath)

	fullS3Path := os.Getenv("S3_DEFAULT_PATH") + "/" + file

	_, err = awsS3Service.GetFile(fullS3Path)

	if nil != err {
		t.Errorf("Failed to get path %s : %s", fullS3Path, err.Error())
	}

}

func TestAWSS3Service_DeleteFile(t *testing.T) {

	pwd, _ := os.Getwd()
	godotenv.Load(pwd + "/../.env")

	var err error
	awsS3Service, err = awsS3Service.Initialize()

	if nil != err {
		t.Errorf("Failed to initialize aws s3 service: %s", err.Error())
	}

	filePath := pwd + "/../" + os.Getenv("TEST_IMAGE_FILE")
	_, file := filepath.Split(filePath)

	fullS3Path := os.Getenv("S3_DEFAULT_PATH") + "/" + file

	_, err = awsS3Service.DeleteFile(fullS3Path)

	if nil != err {
		t.Errorf("Failed to delete path %s : %s", fullS3Path, err.Error())
	}

}
