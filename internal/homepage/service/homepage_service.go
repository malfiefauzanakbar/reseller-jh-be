package service

import (		
	"reseller-jh-be/internal/homepage/model"
	"reseller-jh-be/internal/homepage/repository"		
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

func (s *HomepageService) UpdateHomepage(reqHomepage *model.Homepage) (homepage *model.Homepage, err error) {			
	homepage, err = s.Repo.GetHomepage()
	if err != nil {
		return nil, err
	}

	homepage.BannerTitle         = reqHomepage.BannerTitle
	homepage.BannerImage         = reqHomepage.BannerImage
	homepage.ShortDescription    = reqHomepage.ShortDescription
	homepage.JourneyTitle        = reqHomepage.JourneyTitle
	homepage.JourneyDescription  = reqHomepage.JourneyDescription
	homepage.StoreTitle          = reqHomepage.StoreTitle
	homepage.StoreDescription    = reqHomepage.StoreDescription
	homepage.TierTitle           = reqHomepage.TierTitle
	homepage.TierDescription     = reqHomepage.TierDescription
	homepage.FacilityTitle       = reqHomepage.FacilityTitle
	homepage.FacilityDescription = reqHomepage.FacilityDescription
	homepage.VideoTitle          = reqHomepage.VideoTitle
	homepage.VideoDescription    = reqHomepage.VideoDescription
	homepage.VideoLink           = reqHomepage.VideoLink	
	err = s.Repo.UpdateHomepage(homepage)
	if err != nil {
		return nil, err
	}

	return homepage, nil
}