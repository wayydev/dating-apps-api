package repositories

import (
	"dating-apps/api/cmd/models"
	"dating-apps/api/cmd/repositories/interfaces"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepositoryInterface {
	return &UserRepository{db}
}

func (r UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}