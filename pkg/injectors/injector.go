package injectors

import (
	"dating-apps/api/cmd/repositories"
	"dating-apps/api/cmd/services"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InitService(db *gorm.DB, validate *validator.Validate) *services.Service {
	repo := repositories.NewInitRepository(db)
	service := services.NewInitService(repo, validate)
	return service
}
