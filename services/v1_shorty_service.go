package services

import (
	// "fmt"
	"time"
	"../objects"
	"../repositories"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type V1ShortyService struct {
	request        objects.V1ShortyObjectResponse
	shortyRepository repositories.V1ShortyRepository
}

func V1ShortyServiceHandler(db *gorm.DB) (V1ShortyService) {
	service := V1ShortyService{
		shortyRepository: repositories.V1ShortyRepositoryHandler(db),
	}
	return service
}

func (service *V1ShortyService) GetByShortcode(shortcode string) (objects.V1ShortyObjectResponse, error) {
	shorty, err := service.shortyRepository.GetByShortcode(shortcode)
	if nil != err {
		return objects.V1ShortyObjectResponse{}, err
	}
	shorty.LastSeenDate = time.Now()
	shorty.RedirectCount += 1

	service.shortyRepository.UpdateRedirectCount(shortcode, shorty)
	
	result := objects.V1ShortyObjectResponse{}
	copier.Copy(&result, &shorty)

	return result, nil
}

func (service *V1ShortyService) GetByShortcodeStats(shortcode string) (interface{}, error) {
	shorty, err := service.shortyRepository.GetByShortcodeStats(shortcode)
	if nil != err {
		return objects.V1ShortyObjectResponse{}, err
	}

	// result := objects.V1ShortyObjectResponse{}
	if (shorty.RedirectCount == 0) {
		result := objects.V1ShortyObjectResponseNoLastSeen{}
		copier.Copy(&result, &shorty)
		return result, nil	
	} else {
		result := objects.V1ShortyObjectResponseLastSeen{}
		copier.Copy(&result, &shorty)
		return result, nil	
	}
	
}
