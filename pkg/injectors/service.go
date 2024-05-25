package injectors

import (
	"dating-apps/api/cmd/services"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	AuthService *services.AuthService
}

func NewInitService(repo *Repository, validate *validator.Validate) *Service {
	return &Service{
		AuthService: services.NewAuthService(repo.UserRepository, validate),
	}
}
