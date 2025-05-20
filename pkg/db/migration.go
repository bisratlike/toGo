package db

import "github.com/bisratlike/toGo/internal/auth_service/models"
func RunMigrations() error {
    return DB.AutoMigrate(
        &models.User{},
    )
}