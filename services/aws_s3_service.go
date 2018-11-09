package services

import (
	"../constants"
	"../helpers"
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"net/http"
	"os"
)

type AWSS3Service struct {
	bucketName    string
	errorHandling helpers.ErrorHandling
	credential    *credentials.Credentials
	config        *aws.Config
	s3            *s3.S3
}

func (service *AWSS3Service) Initialize() (AWSS3Service, error) {

	service.bucketName = os.Getenv("S3_BUCKET")
	service.errorHandling = helpers.ErrorHandling{}
	service.credential = credentials.NewStaticCredentials(os.Getenv("S3_KEY"), os.Getenv("S3_SECRET"), "")

	if _, err := service.credential.Get(); nil != err {
		return AWSS3Service{}, err
	}

	service.config = aws.NewConfig().WithRegion(os.Getenv("S3_REGION")).WithCredentials(service.credential)
	service.s3 = s3.New(session.New(), service.config)

	return *service, nil

}

func (service *AWSS3Service) UploadFile(file io.Reader, savePath string, fileName string) (*s3.PutObjectOutput, error) {

	if nil == service.credential {
		return nil, &helpers.CustomError{E: constants.GetErrorConstant(constants.ObjectNotInitializedProperly).Message}
	}

	fileSize := service.getSize(file)

	buffer := make([]byte, fileSize)
	file.Read(buffer)

	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	params := &s3.PutObjectInput{
		Bucket:        aws.String(service.bucketName),
		Key:           aws.String(savePath + "/" + fileName),
		Body:          fileBytes,
		ContentLength: aws.Int64(int64(fileSize)),
		ContentType:   aws.String(fileType),
	}

	return service.s3.PutObject(params)

}

func (service *AWSS3Service) DeleteFile(filePath string) (*s3.DeleteObjectOutput, error) {

	if nil == service.credential {
		return nil, &helpers.CustomError{E: constants.GetErrorConstant(constants.ObjectNotInitializedProperly).Message}
	}

	params := &s3.DeleteObjectInput{
		Bucket: aws.String(service.bucketName),
		Key:    aws.String(filePath),
	}

	return service.s3.DeleteObject(params)

}

func (service *AWSS3Service) GetFile(filePath string) (*s3.GetObjectOutput, error) {

	if nil == service.credential {
		return nil, &helpers.CustomError{E: constants.GetErrorConstant(constants.ObjectNotInitializedProperly).Message}
	}

	params := &s3.GetObjectInput{
		Bucket: aws.String(service.bucketName),
		Key:    aws.String(filePath),
	}

	return service.s3.GetObject(params)

}

func (service *AWSS3Service) getSize(stream io.Reader) int {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Len()
}
