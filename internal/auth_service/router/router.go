package router

import (
    "github.com/go-chi/chi/v5"
    "gorm.io/gorm"
    "github.com/bisratlike/toGo/internal/auth_service/handler"
    "github.com/bisratlike/toGo/internal/auth_service/service"
)

func AuthRoutes(r chi.Router, db *gorm.DB) {
    authService := &service.AuthService{DB: db}
    authHandler := handler.NewAuthHandler(authService)

    r.Route("/api/auth", func(r chi.Router) {
        r.Post("/register", authHandler.RegisterHandler)
        // r.Post("/login", authHandler.LoginHandler)
        // r.Post("/logout", authHandler.LogoutHandler)
    })
}