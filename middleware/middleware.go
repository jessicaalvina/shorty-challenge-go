package middleware

import "github.com/jinzhu/gorm"

type DefaultMiddleware struct {
	DB *gorm.DB
}
