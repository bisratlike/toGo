
package handler

import (
    "net/http"
    "github.com/bisratlike/toGo/internal/auth_service/dto"
	"github.com/bisratlike/toGo/internal/auth_service/service"
    "github.com/bisratlike/toGo/pkg/validator"
    "github.com/bisratlike/toGo/pkg/response"
	"github.com/bisratlike/toGo/pkg/security"
    // "github.com/go-playground/validator/v10"
)


type AuthHandler struct {
    service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
    return &AuthHandler{
        service: service,
    }
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var req dto.RegisterRequest

    if ok := validator.ParseAndValidate(w, r, &req); !ok {
        return
    }

    user, err := h.service.Register(req)
    if err != nil {
        if err == service.ErrEmailAlreadyInUse {
            response.Error(w, http.StatusBadRequest, "Email already in use", err)
            return
        }
        response.Error(w, http.StatusInternalServerError, "Failed to register user", err)
        return
    }

    token, err := security.GenerateJWT(user.ID, user.Role)
    if err != nil {
        response.Error(w, http.StatusInternalServerError, "Failed to generate token", err)
        return
    }

    response.Success(w, http.StatusCreated, "User registered successfully", map[string]interface{}{
        "user":  user,
        "token": token,
    })
}


// func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
//     var req dto.LoginRequest
//     if ok := ParseAndValidate(w, r, &req); !ok {
//         return
//     }

//     user, err := h.service.Login(req)
//     if err != nil {
//         if err == service.ErrInvalidCredentials {
//             response.Error(w, http.StatusUnauthorized, "Invalid credentials")
//         } else {
//             response.Error(w, http.StatusInternalServerError, "Could not login")
//         }
//         return
//     }

//     token, err := generateToken(user.ID)
//     if err != nil {
//         response.Error(w, http.StatusInternalServerError, "Could not generate token")
//         return
//     }

//     resp := map[string]interface{}{
//         "user":  user,
//         "token": token,
//     }
//     response.Success(w, http.StatusOK, "Login successful", resp)
// }
