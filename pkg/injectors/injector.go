package injectors

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InitService(db *gorm.DB, validate *validator.Validate) *Service {
	repo := NewInitRepository(db)
	service := NewInitService(repo, validate)
	return service
}
