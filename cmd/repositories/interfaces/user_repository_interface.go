package interfaces

import "dating-apps/api/cmd/models"

type UserRepositoryInterface interface {
	FindById(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
}
