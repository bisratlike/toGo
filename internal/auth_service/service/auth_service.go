package service

import (
    "errors"
    "strings"
    "github.com/bisratlike/toGo/internal/auth_service/dto"
    "github.com/bisratlike/toGo/internal/auth_service/models"
    "github.com/bisratlike/toGo/internal/auth_service/repository"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

var ErrEmailAlreadyInUse = errors.New("email already in use")

type AuthService struct {
    DB *gorm.DB
}

func (s *AuthService) Register(req dto.RegisterRequest) (*models.User, error) {
    hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := models.User{
        FullName: req.FullName,
        Email:    req.Email,
        Password: string(hashedPwd),
    }

    createdUser, err := repository.CreateUser(s.DB, user)
    if err != nil {
        if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "UNIQUE") {
            return nil, ErrEmailAlreadyInUse
        }
        return nil, err
    }
    return createdUser, nil
}