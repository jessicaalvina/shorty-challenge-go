package models

import (
	"time"
)

type Order struct {
	ID          uint       `gorm:"column:id;primary_key" json:"id"`
	UserId      uint       `gorm:"column:user_id" json:"user_id"`
	OrderSerial string     `gorm:"column:order_serial" json:"order_serial"`
	GrandTotal  float64    `gorm:"column:grand_total" json:"grand_total"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

func (Order) TableName() string {
	return "rl_orders"
}
