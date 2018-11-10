package repositories

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"ralali.com/models"
	"ralali.com/objects"
)

type UserRepository struct {
	DB gorm.DB
}

func UserRepositoryHandler(db *gorm.DB) (UserRepository) {
	repository := UserRepository{DB: *db}
	return repository
}

func (repository *UserRepository) GetById(id int) (models.User, error) {

	userResponse := models.User{}

	query := repository.DB.Table("rl_users")
	query = query.Where("id=?", id)
	query = query.First(&userResponse)

	return userResponse, query.Error

}

func (repository *UserRepository) UpdateById(id int, userData objects.UserObject) (models.User, error) {

	userModel := models.User{}
	copier.Copy(&userModel, &userData)

	fmt.Println(userModel)
	fmt.Println(userData)

	query := repository.DB.Table("rl_users")
	query = query.Where("id=?", id)
	query = query.Updates(userModel)
	query = query.Scan(&userModel)

	return userModel, query.Error

}
