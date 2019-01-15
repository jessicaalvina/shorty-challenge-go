package repositories

import (
	"time"

	"../models"
	"../objects"

	// "github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type V2ShortyRepository struct {
	DB gorm.DB
}

func V2ShortyRepositoryHandler(db *gorm.DB) V2ShortyRepository {
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
