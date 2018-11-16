package services

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"ralali.com/objects"
	"ralali.com/repositories"
)

type V1UserService struct {
	request        objects.V1UserObjectResponse
	userRepository repositories.V1UserRepository
}

func UserServiceHandler(db *gorm.DB) (V1UserService) {
	service := V1UserService{
		userRepository: repositories.V1UserRepositoryHandler(db),
	}
	return service
}

func (service *V1UserService) GetById(id int) (objects.V1UserObjectResponse, error) {
	user, err := service.userRepository.GetById(id)
	if nil != err {
		return objects.V1UserObjectResponse{}, err
	}
	result := objects.V1UserObjectResponse{}
	copier.Copy(&result, &user)
	return result, nil
}

func (service *V1UserService) UpdateById(id int, requestObject objects.V1UserObjectRequest) (objects.V1UserObjectResponse, error) {

	fmt.Println(requestObject)

	user, err := service.userRepository.UpdateById(id, requestObject)
	if nil != err {
		return objects.V1UserObjectResponse{}, err
	}

	result := objects.V1UserObjectResponse{}
	copier.Copy(&result, &user)

	return result, nil

}
