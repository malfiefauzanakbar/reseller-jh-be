package repository

import (	
	"reseller-jh-be/internal/homepage/model"	

	"gorm.io/gorm"
)

type HomepageRepository struct {
	DB *gorm.DB
}

type HomepageRepositoryInterface interface {
}

func NewHomepageRepository(DB *gorm.DB) *HomepageRepository {
	return &HomepageRepository{
		DB: DB,
	}
}

func (r *HomepageRepository) GetHomepage() (*model.Homepage, error) {
	var homepage model.Homepage
	if err := r.DB.First(&homepage).Error; err != nil {
		return nil, err
	}
	return &homepage, nil
}

func (r *HomepageRepository) UpdateHomepage(homepage *model.Homepage) error {
	return r.DB.Save(homepage).Error
}