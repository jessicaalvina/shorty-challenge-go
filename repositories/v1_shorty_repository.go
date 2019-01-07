package repositories

import (
	// "fmt"
	"../models"
	"github.com/jinzhu/gorm"
)

type V1ShortyRepository struct {
	DB gorm.DB
}

func V1ShortyRepositoryHandler(db *gorm.DB) (V1ShortyRepository) {
	repository := V1ShortyRepository{DB: *db}
	return repository
}

func (repository *V1ShortyRepository) UpdateRedirectCount(shortcode string, shorty interface{}) (bool){
	query := repository.DB.Table("shorty")
	query = query.Where("shortcode = ?", shortcode)
	query = query.Updates(shorty)
	
	return true
}

func (repository *V1ShortyRepository) GetByShortcode(shortcode string) (models.Shorty, error) {
	shortyResponse := models.Shorty{}

	query := repository.DB.Table("shorty")
	query = query.Where("shortcode = ?", shortcode)
	query = query.First(&shortyResponse)
	
	return shortyResponse, query.Error
}

func (repository *V1ShortyRepository) GetByShortcodeStats(shortcode string) (models.Shorty, error) {
	shortyResponse := models.Shorty{}
	
	query := repository.DB.Table("shorty")
	query = query.Where("shortcode = ?", shortcode)
	query = query.First(&shortyResponse)

	return shortyResponse, query.Error
}