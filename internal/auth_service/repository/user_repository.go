package repository

import (
	"github.com/bisratlike/toGo/internal/auth_service/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user models.User) (*models.User, error) {
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}