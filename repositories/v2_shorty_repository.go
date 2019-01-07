package repositories

import (
	"time"
	"regexp"
	"net/url"
	"../objects"
	"../models"
	// "github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type V2ShortyRepository struct {
	DB gorm.DB
}

func V2ShortyRepositoryHandler(db *gorm.DB) (V2ShortyRepository) {
	repository := V2ShortyRepository{DB: *db}
	return repository
}

func (repository *V2ShortyRepository) PostByShorten(requestObject objects.V2ShortyObjectRequest) (models.Shorty, error) {
	shortyModel := models.Shorty{}

	shortyModel.Url = requestObject.Url
	shortyModel.Shortcode = requestObject.Shortcode
	shortyModel.StartDate = time.Now()
	shortyModel.LastSeenDate = time.Now()
	shortyModel.RedirectCount = 0

	query := repository.DB.Table("shorty")
	query = query.Create(&shortyModel)
	
	return shortyModel, query.Error
}

func (repository *V2ShortyRepository) ValidateShortcode(shortcode string) (bool) {
	validate,_ := regexp.MatchString("^[0-9a-zA-Z_]{6}$", shortcode)

	return validate
}

func (repository *V2ShortyRepository) ValidateUrl(urlInput string) (bool) {
	_, err := url.ParseRequestURI(urlInput)
    if err != nil {
        return false
    } else {
        return true
    }
}
