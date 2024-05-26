package services

import (
	"dating-apps/api/cmd/models"
	"dating-apps/api/cmd/repositories/interfaces"
	"dating-apps/api/cmd/requests"
	"dating-apps/api/pkg/utilities"

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

	if err := utilities.HashCompare(request.Password, user.Password); err != false {
		return nil, utilities.Error("Username or password is wrong", 422, nil, nil)
	}

	return user, nil
}

func (s AuthService) Registration() {

}
