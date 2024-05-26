package databases

import (
	"dating-apps/api/cmd/models"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Package{},
		&models.PackageFeature{},
		&models.User{},
		&models.UserPackage{},
		&models.UserPhoto{},
		&models.Notification{},
		&models.Swap{},
	)
}
