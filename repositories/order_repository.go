package repositories

import (
	"../models"
	"../requests"
	"github.com/jinzhu/gorm"
)

type OrderRepository struct {
	DB      gorm.DB
	request requests.UserRequest
}

func (repository *OrderRepository) GetListByUserId(userId int, page int, perPage int) ([]models.Order, error) {
	var orderResponse []models.Order

	query := repository.DB.Table("rl_orders")
	query = query.Where("user_id=?", userId)
	query = query.Offset((page - 1) * perPage).Limit(perPage)
	query = query.Scan(&orderResponse)

	return orderResponse, query.Error
}
