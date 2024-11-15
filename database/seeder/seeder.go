// /database/migration/seeder.go

package migration

import (
	homepage "reseller-jh-be/internal/homepage/model"
	status "reseller-jh-be/internal/status/model"
	user "reseller-jh-be/internal/user/model"
	"reseller-jh-be/pkg/common"

	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) error {

	var existingUser user.User
	password := "admin123"
	hashedPassword, err := common.HashPassword(password)
	if err != nil {
		return err
	}
	dataUsers := []user.User{
		{Username: "admin", Password: hashedPassword, Email: "admin@gmail.com"},
	}
	for _, user := range dataUsers {
		err := db.Where("username = ?", user.Username).First(&existingUser).Error
		if err == gorm.ErrRecordNotFound {
			if err = db.Create(&user).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func SeedStatus(db *gorm.DB) error {
	dataStatuss := []status.Status{
		{Name: "Unread"},
		{Name: "Read"},
		{Name: "Process"},
		{Name: "Discuss"},
		{Name: "Cancel"},
		{Name: "Deal"},
	}

	var existingStatus status.Status
	for _, status := range dataStatuss {
		err := db.Where("name = ?", status.Name).First(&existingStatus).Error
		if err == gorm.ErrRecordNotFound {
			if err = db.Create(&status).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func SeedHomepage(db *gorm.DB) error {
	dataHomepages := []homepage.Homepage{
		{
			BannerTitle:         "Banner Title",
			BannerImage:         "BannerImage.jpg",
			ShortDescription:    "Short Description",
			JourneyTitle:        "Journey Title",
			JourneyDescription:  "Journey Description",
			StoreTitle:          "Store Title",
			StoreDescription:    "Store Description",
			TierTitle:           "Tier Title",
			TierDescription:     "Tier Description",
			FacilityTitle:       "Facility Title",
			FacilityDescription: "Facility Description",
			VideoTitle:          "Video Title",
			VideoDescription:    "Video Description",
			VideoLink:           "https://youtube.com",
		},
	}

	var existingHomepage homepage.Homepage
	for _, homepage := range dataHomepages {
		err := db.Where("banner_title = ?", homepage.BannerTitle).First(&existingHomepage).Error
		if err == gorm.ErrRecordNotFound {
			if err = db.Create(&homepage).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
