package services

import (
	"../objects"
	"../repositories"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type V2ShortyService struct {
	request        objects.V2ShortyObjectResponse
	shortyRepository repositories.V2ShortyRepository
}

func V2ShortyServiceHandler(db *gorm.DB) (V2ShortyService) {
	service := V2ShortyService{
		shortyRepository: repositories.V2ShortyRepositoryHandler(db),
	}
	return service
}

func (service *V2ShortyService) ValidateShortcode(shortcode string) (bool) {
	validate := service.shortyRepository.ValidateShortcode(shortcode)
	return validate
}

func (service *V2ShortyService) ValidateUrl(url string) (bool) {
	validate := service.shortyRepository.ValidateUrl(url)
	return validate
}

func (service *V2ShortyService) PostByShorten(requestObject objects.V2ShortyObjectRequest) (objects.V2ShortyObjectResponse, error) {
	shorty, err := service.shortyRepository.PostByShorten(requestObject)
	if nil != err {
		return objects.V2ShortyObjectResponse{}, err
	}

	result := objects.V2ShortyObjectResponse{}
	copier.Copy(&result, &shorty)

	return result, nil
}
