package services

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"ralali.com/objects"
	"ralali.com/repositories"
)

type UserService struct {
	request        objects.UserObject
	userRepository repositories.UserRepository
}

func UserServiceHandler(db *gorm.DB) (UserService) {
	service := UserService{
		userRepository: repositories.UserRepositoryHandler(db),
	}
	return service
}

func (service *UserService) GetById(id int) (objects.UserObject, error) {
	user, err := service.userRepository.GetById(id)
	if nil != err {
		return objects.UserObject{}, err
	}
	result := objects.UserObject{}
	copier.Copy(&result, &user)
	return result, nil
}

func (service *UserService) UpdateById(id int, requestObject objects.UserObject) (objects.UserObject, error) {

	fmt.Println(requestObject)

	user, err := service.userRepository.UpdateById(id, requestObject)
	if nil != err {
		return objects.UserObject{}, err
	}

	result := objects.UserObject{}
	copier.Copy(&result, &user)

	return result, nil

}
