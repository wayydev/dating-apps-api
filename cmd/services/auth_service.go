package services

import (
	"dating-apps/api/cmd/models"
	"dating-apps/api/cmd/repositories"
	"dating-apps/api/cmd/requests"
	"dating-apps/api/cmd/responses"
	"dating-apps/api/pkg/utilities"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type AuthService struct {
	Repo     *repositories.Repository
	Validate *validator.Validate
}

func NewAuthService(repo *repositories.Repository, validate *validator.Validate) *AuthService {
	return &AuthService{
		Repo:     repo,
		Validate: validate,
	}
}

func (s AuthService) Login(request *requests.Login) (*responses.Login, error) {
	if err := s.Validate.Struct(request); err != nil {
		return nil, err
	}

	user, err := s.Repo.UserRepository.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	if hash := utilities.HashCompare(request.Password, user.Password); !hash {
		return nil, utilities.Error("Username or password is wrong", 422, nil, nil)
	}

	token, err := utilities.CreateTokenJWT(&utilities.JWT{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	})

	if err != nil {
		return nil, utilities.Error("Failed to create auth token", 500, nil, err.Error())
	}

	response := &responses.Login{
		Token: token,
		User:  user,
	}

	return response, nil
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
		Location:       fmt.Sprintf("SRID=4326;POINT(%f %f)", request.Longitude, request.Latitude),
		PlaceOfBirth:   request.PlaceOfBirth,
		DateOfBirth:    request.DateOfBirth,
		FindOnDistance: request.FindOnDistance,
		Gender:         request.Gender,
	}

	_, err = s.Repo.UserRepository.Create(user)

	if err != nil {
		return utilities.Error("Failed to create a user", 500, nil, err)
	}

	return nil
}
