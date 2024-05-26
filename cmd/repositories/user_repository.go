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

func (r UserRepository) FindById(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).Find(&user).Error; err != nil {
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

func (r UserRepository) Create(user *models.User) (*models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	if err := r.db.Create(&models.UserPackage{
		UserId:    user.ID,
		PackageId: 1,
	}).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepository) GetAllNotification(id uint) (*[]models.Notification, error) {
	var user []models.Notification
	if err := r.db.Where("user_id", id).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
