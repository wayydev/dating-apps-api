package services

import (
	"dating-apps/api/cmd/models"
	"dating-apps/api/cmd/repositories/interfaces"
	"dating-apps/api/cmd/requests"

	"github.com/go-playground/validator/v10"
)

type AuthService struct {
	UserRepository interfaces.UserRepositoryInterface
	Validate       *validator.Validate
}

func NewAuthService(userRepository interfaces.UserRepositoryInterface, validate *validator.Validate) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (s AuthService) Login(request *requests.Login) (*models.User, error) {
	if err := s.Validate.Struct(request); err != nil {
		return nil, err
	}

	user, err := s.UserRepository.FindByUsername(request.Username)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s AuthService) Registration() {

}
