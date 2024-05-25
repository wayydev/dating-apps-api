package interfaces

import "dating-apps/api/cmd/models"

type UserRepositoryInterface interface {
	FindByEmail(email string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
}
