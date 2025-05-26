package service

import (
	"testing"
	"github.com/bisratlike/toGo/internal/auth_service/dto"
	"github.com/bisratlike/toGo/internal/auth_service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&models.User{})
	return db
}

func TestAuthService_Register(t *testing.T) {
	db := setupTestDB(t)
	service := &AuthService{DB: db}

	t.Run("Success", func(t *testing.T) {
		req := dto.RegisterRequest{
			FullName: "John Doe",
			Email:    "john@example.com",
			Password: "password123",
		}

		user, err := service.Register(req)
		assert.NoError(t, err)
		assert.Equal(t, "John Doe", user.FullName)
		assert.Equal(t, "john@example.com", user.Email)
		
	
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		assert.NoError(t, err)
	})

	t.Run("Duplicate Email", func(t *testing.T) {
		req := dto.RegisterRequest{
			FullName: "Jane Doe",
			Email:    "jane@example.com",
			Password: "password123",
		}

		_, err := service.Register(req)
		assert.NoError(t, err)
		
		_, err = service.Register(req)
		assert.Equal(t, ErrEmailAlreadyInUse, err)
	})
}