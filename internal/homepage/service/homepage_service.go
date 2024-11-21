package service

import (	
	"mime/multipart"
	"reseller-jh-be/internal/homepage/model"
	"reseller-jh-be/internal/homepage/repository"
	"reseller-jh-be/internal/homepage/request"
	"reseller-jh-be/pkg/common"

	"github.com/gin-gonic/gin"
)

type HomepageService struct {
	Repo repository.HomepageRepository
}

type HomepageServiceInterface interface {
}

func NewHomepageService(repo *repository.HomepageRepository) *HomepageService {
	return &HomepageService{
		Repo: *repo,
	}
}

func (s *HomepageService) GetHomepage() (homepage *model.Homepage, err error) {
	homepage, err = s.Repo.GetHomepage()
	if err != nil {
		return nil, err
	}

	return homepage, nil
}

func (s *HomepageService) UpdateHomepage(c *gin.Context, reqHomepage *request.ReqHomepage, file *multipart.FileHeader) (homepage *model.Homepage, err error) {
	var filePath string
	if file != nil {
		filePath, err = common.UploadFile(c, file, "bannerImage")
		if err != nil {
			return nil, err
		}
	}		

	homepage, err = s.Repo.GetHomepage()
	if err != nil {
		return nil, err
	}

	if filePath == "" {		
		filePath = homepage.BannerImage
	}	

	homepage.BannerTitle = reqHomepage.BannerTitle
	homepage.BannerImage = filePath
	homepage.ShortDescription = reqHomepage.ShortDescription
	homepage.JourneyTitle = reqHomepage.JourneyTitle
	homepage.JourneyDescription = reqHomepage.JourneyDescription
	homepage.StoreTitle = reqHomepage.StoreTitle
	homepage.StoreDescription = reqHomepage.StoreDescription
	homepage.TierTitle = reqHomepage.TierTitle
	homepage.TierDescription = reqHomepage.TierDescription
	homepage.FacilityTitle = reqHomepage.FacilityTitle
	homepage.FacilityDescription = reqHomepage.FacilityDescription
	homepage.VideoTitle = reqHomepage.VideoTitle
	homepage.VideoDescription = reqHomepage.VideoDescription
	homepage.VideoLink = reqHomepage.VideoLink	
	err = s.Repo.UpdateHomepage(homepage)
	if err != nil {
		return nil, err
	}

	return homepage, nil
}
