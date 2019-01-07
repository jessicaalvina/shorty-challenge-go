package models

import (
	"time"
)

type Shorty struct {
	ID           	uint       `gorm:"column:id;primary_key" json:"id"`
	Url          	string     `gorm:"column:url" json:"url"`
	Shortcode    	string     `gorm:"column:shortcode" json:"shortcode"`
	RedirectCount   int		   `gorm:"column:redirect_count" json:"redirect_count"`
	StartDate    	time.Time  `gorm:"column:start_date" json:"start_date"`
	LastSeenDate    time.Time  `gorm:"column:last_seen_date" json:"last_seen_date"`
}

func (Shorty) TableName() string {
	return "shorty"
}
