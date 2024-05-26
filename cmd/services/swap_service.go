package services

import (
	"dating-apps/api/cmd/models"
	"dating-apps/api/cmd/repositories"
	"dating-apps/api/pkg/utilities"

	"github.com/go-playground/validator/v10"
)

type SwapService struct {
	Repo     *repositories.Repository
	Validate *validator.Validate
}

func NewSwapService(repo *repositories.Repository, validate *validator.Validate) *SwapService {
	return &SwapService{
		Repo:     repo,
		Validate: validate,
	}
}

func (s SwapService) Find(userID uint) (*models.User, error) {
	user, err := s.Repo.SwapRepository.FindUser(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s SwapService) Like(userAuth *utilities.JWT, swapUserID uint) (*models.User, error) {
	if err := s.Repo.SwapRepository.Like(userAuth, swapUserID); err != nil {
		return nil, err
	}

	user, err := s.Repo.SwapRepository.FindUser(userAuth.ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s SwapService) Pass(userID uint, swapUserID uint) (*models.User, error) {
	if err := s.Repo.SwapRepository.Pass(userID, swapUserID); err != nil {
		return nil, err
	}

	user, err := s.Repo.SwapRepository.FindUser(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s SwapService) Limiter(userID uint) error {
	err := s.Repo.SwapRepository.Limiter(userID)

	if err != nil {
		return err
	}

	return nil
}
