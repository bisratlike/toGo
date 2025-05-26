package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/bisratlike/toGo/internal/auth_service/dto"
	"github.com/bisratlike/toGo/internal/auth_service/models"
	"github.com/bisratlike/toGo/internal/auth_service/service"
	"github.com/bisratlike/toGo/pkg/response"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type MockAuthService struct {
	registerFunc func(dto.RegisterRequest) (*models.User, error)
}

func (m *MockAuthService) Register(req dto.RegisterRequest) (*models.User, error) {
	return m.registerFunc(req)
}

func TestRegisterHandler_Success(t *testing.T) {
	mockService := &MockAuthService{
		registerFunc: func(req dto.RegisterRequest) (*models.User, error) {
			return &models.User{
				ID:       uuid.New(),
				FullName: req.FullName,
				Email:    req.Email,
				Role:     "user",
			}, nil
		},
	}

	handler := NewAuthHandler(mockService)

	payload := []byte(`{
		"full_name": "Test User",
		"email": "test@example.com",
		"password": "password123"
	}`)

	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.RegisterHandler(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	
	var resp response.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.True(t, resp.Success)
}

func TestRegisterHandler_EmailConflict(t *testing.T) {
	mockService := &MockAuthService{
		registerFunc: func(req dto.RegisterRequest) (*models.User, error) {
			return nil, service.ErrEmailAlreadyInUse
		},
	}

	handler := NewAuthHandler(mockService)

	payload := []byte(`{
		"full_name": "Test User",
		"email": "exists@example.com",
		"password": "password123"
	}`)

	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.RegisterHandler(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var resp response.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.False(t, resp.Success)
	assert.Contains(t, resp.Message, "Email already in use")
}