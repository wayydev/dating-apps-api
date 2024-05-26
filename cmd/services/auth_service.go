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

	if hash := utilities.HashCompare(request.Password, user.Password); !hash {
		return nil, utilities.Error("Username or password is wrong", 422, nil, nil)
	}

	return user, nil
}

func (s AuthService) Registration(request *requests.Registration) error {
	if err := s.Validate.Struct(request); err != nil {
		return err
	}

	password, err := utilities.Hash(request.Password)

	if err != nil {
		return utilities.Error("Failed to make a password", 422, nil, err)
	}

	user := &models.User{
		Name:           request.Name,
		Username:       request.Username,
		Email:          request.Email,
		Phone:          request.Phone,
		Password:       password,
		Latitude:       request.Latitude,
		Longitude:      request.Longitude,
		PlaceOfBirth:   request.PlaceOfBirth,
		DateOfBirth:    request.DateOfBirth,
		PhotoProfile:   request.PhotoProfile,
		FindOnDistance: request.FindOnDistance,
	}

	_, err = s.UserRepository.Create(user)

	if err != nil {
		return utilities.Error("Failed to create a user", 500, nil, err)
	}

	return nil
}
