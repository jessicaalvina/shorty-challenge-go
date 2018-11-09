package repositories

import (
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"ralali.com/models"
	"ralali.com/requests"
)

type UserRepository struct {
	DB      gorm.DB
	request requests.UserRequest
}

func (repository *UserRepository) GetById(id int) (models.User, error) {
	userResponse := models.User{}
	query := repository.DB.Table("rl_users").Where("id=?", id).First(&userResponse)
	return userResponse, query.Error
}

func (repository *UserRepository) GetList(page int, perPage int) ([]models.User, error) {
	var userResponse []models.User
	query := repository.DB.Table("rl_users").Offset((page - 1) * perPage).Limit(perPage).Scan(&userResponse)
	return userResponse, query.Error
}

func (repository *UserRepository) UpdateUser(userId int, userRequest interface{}) (models.User, error) {
	user := models.User{}
	copier.Copy(&user, &userRequest)
	query := repository.DB.Table("rl_users").Where("id=?", userId).Omit("created_at", "deleted_at", "id").Updates(user).Scan(&user)
	return user, query.Error
}
