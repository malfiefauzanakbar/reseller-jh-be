// /database/migration/migrate.go

package migration

import (
	seed "reseller-jh-be/database/seeder"
	homepage "reseller-jh-be/internal/homepage/model"
	reseller "reseller-jh-be/internal/reseller/model"
	status "reseller-jh-be/internal/status/model"
	user "reseller-jh-be/internal/user/model"
	"reseller-jh-be/pkg/common"

	"gorm.io/gorm"
)

// MigrateDB handles the database migrations
func MigrateDB(db *gorm.DB) {
	common.Log.Info("===== MIGRATE TABLE - user =====")
	if err := db.AutoMigrate(&user.User{}); err != nil {
		common.Log.Error(err)
	}

	common.Log.Info("===== MIGRATE TABLE - resellers =====")
	if err := db.AutoMigrate(&reseller.Reseller{}); err != nil {
		common.Log.Error(err)
	} else {
		err = db.Exec("ALTER TABLE resellers ALTER COLUMN nik TYPE VARCHAR(50);").Error
		if err != nil {
			common.Log.Error(err)
		}
	}

	common.Log.Info("===== MIGRATE TABLE - status =====")
	if err := db.AutoMigrate(&status.Status{}); err != nil {
		common.Log.Error(err)
	}

	common.Log.Info("===== MIGRATE TABLE - homepages =====")
	if err := db.AutoMigrate(&homepage.Homepage{}); err != nil {
		common.Log.Error(err)
	}

	//Seeder
	common.Log.Info("===== SEED TABLE - users =====")
	if err := seed.SeedUser(db); err != nil {
		common.Log.Error(err)
	}

	common.Log.Info("===== SEED TABLE - status =====")
	if err := seed.SeedStatus(db); err != nil {
		common.Log.Error(err)
	}

	common.Log.Info("===== SEED TABLE - homepages =====")
	if err := seed.SeedHomepage(db); err != nil {
		common.Log.Error(err)
	}
}
