package services

import (
	"dating-apps/api/cmd/repositories"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	AuthService    *AuthService
	ProfileService *ProfileService
	SwapService    *SwapService
}

func NewInitService(repo *repositories.Repository, validate *validator.Validate) *Service {
	return &Service{
		AuthService:    NewAuthService(repo, validate),
		ProfileService: NewProfileService(repo, validate),
		SwapService:    NewSwapService(repo, validate),
	}
}
