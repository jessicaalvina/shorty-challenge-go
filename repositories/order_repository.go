package repositories

import (
	"github.com/jinzhu/gorm"
	"ralali.com/models"
	"ralali.com/objects"
)

type OrderRepository struct {
	DB      gorm.DB
	request objects.UserObject
}

func (repository *OrderRepository) GetListByUserId(userId int, page int, perPage int) ([]models.Order, error) {
	var orderResponse []models.Order

	query := repository.DB.Table("rl_orders")
	query = query.Where("user_id=?", userId)
	query = query.Offset((page - 1) * perPage).Limit(perPage)
	query = query.Scan(&orderResponse)

	return orderResponse, query.Error
}
