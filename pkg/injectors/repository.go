package injectors

import (
	"dating-apps/api/cmd/repositories"
	"dating-apps/api/cmd/repositories/interfaces"

	"gorm.io/gorm"
)

type Repository struct {
	UserRepository interfaces.UserRepositoryInterface
}

func NewInitRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: repositories.NewUserRepository(db),
	}
}
