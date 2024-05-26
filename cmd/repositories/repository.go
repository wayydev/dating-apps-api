package repositories

import (
	"dating-apps/api/cmd/repositories/interfaces"

	"gorm.io/gorm"
)

type Repository struct {
	UserRepository interfaces.UserRepositoryInterface
	SwapRepository interfaces.SwapRepositoryInterface
}

func NewInitRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
		SwapRepository: NewSwapRepository(db),
	}
}
