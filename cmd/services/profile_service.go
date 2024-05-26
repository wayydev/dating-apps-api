package services

import (
	"dating-apps/api/cmd/models"
	"dating-apps/api/cmd/repositories"
	"dating-apps/api/pkg/utilities"

	"github.com/go-playground/validator/v10"
)

type ProfileService struct {
	Repo     *repositories.Repository
	Validate *validator.Validate
}

func NewProfileService(repo *repositories.Repository, validate *validator.Validate) *ProfileService {
	return &ProfileService{
		Repo:     repo,
		Validate: validate,
	}
}

func (s ProfileService) GetProfile(id uint) (*models.User, error) {
	user, err := s.Repo.UserRepository.FindById(id)

	if err != nil && user == nil {
		return nil, utilities.Error("User not found", 404, nil, err)
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
