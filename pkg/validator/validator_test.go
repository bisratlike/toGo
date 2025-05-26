package validator

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/bisratlike/toGo/internal/auth_service/dto"
	"github.com/stretchr/testify/assert"
)

func TestParseAndValidate(t *testing.T) {
	t.Run("Invalid JSON", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/", strings.NewReader(`{"invalid`))
		w := httptest.NewRecorder()
		var data dto.RegisterRequest
		
		ok := ParseAndValidate(w, req, &data)
		assert.False(t, ok)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Validation Errors", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/", strings.NewReader(`{"email":"invalid"}`))
		w := httptest.NewRecorder()
		var data dto.RegisterRequest
		
		ok := ParseAndValidate(w, req, &data)
		assert.False(t, ok)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}