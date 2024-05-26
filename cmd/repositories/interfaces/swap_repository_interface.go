package interfaces

import (
	"dating-apps/api/cmd/models"
	"dating-apps/api/pkg/utilities"
)

type SwapRepositoryInterface interface {
	FindUser(userID uint) (*models.User, error)
	Like(user *utilities.JWT, userSwapId uint) error
	Pass(userID uint, userSwapId uint) error
	Limiter(userID uint) error
}
