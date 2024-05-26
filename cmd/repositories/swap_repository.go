package repositories

import (
	"dating-apps/api/cmd/models"
	"dating-apps/api/cmd/repositories/interfaces"
	"dating-apps/api/pkg/utilities"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type SwapRepository struct {
	db *gorm.DB
}

func NewSwapRepository(db *gorm.DB) interfaces.SwapRepositoryInterface {
	return &SwapRepository{db}
}

func (r *SwapRepository) FindUser(userID uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	now := time.Now()
	previous24Hours := now.Add(-24 * time.Hour)

	var swapHistories []models.Swap
	if err := r.db.Where("user_id = ? AND created_at BETWEEN ? AND ?", userID, previous24Hours, now).
		Find(&swapHistories).Error; err != nil {
		return nil, err
	}

	var swapUserIDs []uint
	swapUserIDs = append(swapUserIDs, userID)

	for _, swapHistory := range swapHistories {
		swapUserIDs = append(swapUserIDs, swapHistory.SwapUserID)
	}

	var nearbyUsers []models.User
	if err := r.db.Where("gender != ? AND id NOT IN (?)", user.Gender, swapUserIDs).
		Where("ST_DWithin(ST_MakePoint(?, ?)::geography, ST_MakePoint(longitude, latitude)::geography, ?)", user.Longitude, user.Latitude, float64(user.FindOnDistance)*1000).
		Order("RANDOM()").
		Limit(1).
		Find(&nearbyUsers).Error; err != nil {
		return nil, err
	}

	if len(nearbyUsers) == 0 {
		return nil, utilities.Error("No match result", 404, nil, nil)
	}

	return &nearbyUsers[0], nil
}

func (r *SwapRepository) Like(user *utilities.JWT, swapUserID uint) error {
	swap := models.Swap{
		UserID:     user.ID,
		SwapUserID: swapUserID,
		Like:       true,
	}

	if err := r.db.Create(&swap).Error; err != nil {
		return err
	}

	notification := models.Notification{
		UserID:  swapUserID,
		Title:   fmt.Sprintf("%s Likes You!", user.Name),
		Message: fmt.Sprintf("%s matched with your, contact now!", user.Name),
		Data:    nil,
		IsRead:  false,
	}

	if err := r.db.Create(&notification).Error; err != nil {
		return err
	}

	return nil
}

func (r *SwapRepository) Pass(userID, swapUserID uint) error {
	swap := models.Swap{
		UserID:     userID,
		SwapUserID: swapUserID,
		Like:       false,
	}

	if err := r.db.Create(&swap).Error; err != nil {
		return err
	}

	return nil
}

func (r *SwapRepository) Limiter(userID uint) error {
	var count int64

	now := time.Now()
	previous24Hours := now.Add(-24 * time.Hour)

	if err := r.db.Model(&models.Swap{}).Where("user_id = ? AND created_at BETWEEN ? AND ?", userID, previous24Hours, now).Count(&count).Error; err != nil {
		return err
	}

	var userPackage models.UserPackage
	if err := r.db.Where("user_id = ? AND (expires_at >= ? or expires_at is null)", userID, now).First(&userPackage).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return utilities.Error("User package not found", 404, nil, nil)
		}
		return err
	}

	var packageModel models.Package
	if err := r.db.Where("id = ?", userPackage.PackageID).First(&packageModel).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return utilities.Error("Package not found", 404, nil, nil)
		}
		return err
	}

	var feature models.PackageFeature
	if err := r.db.Where("package_id = ? AND name = ?", packageModel.ID, "limit").First(&feature).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return utilities.Error("Feature not found", 404, nil, nil)
		}
		return err
	}

	limit, err := strconv.Atoi(feature.Value)
	if err != nil {
		return utilities.Error("Invalid feature value", 500, nil, nil)
	}

	if count >= int64(limit) {
		return utilities.Error("Your daily access has expired", 422, nil, nil)
	}

	return nil
}
