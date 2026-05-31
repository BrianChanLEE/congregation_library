package api_test

import (
	"boock/backGo/internal/api"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSystemService struct {
	mock.Mock
}

func (m *MockSystemService) GetSystemErrors() ([]map[string]interface{}, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]map[string]interface{}), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestSystemHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetSystemStatus", func(t *testing.T) {
		mockSvc := new(MockSystemService)
		handler := api.NewSystemHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("GetSystemErrors").Return([]map[string]interface{}{{"err": "none"}}, nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/status", handler.GetSystemStatus)
			req, _ := http.NewRequest("GET", "/status", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("에러", func(t *testing.T) {
			mockSvc.On("GetSystemErrors").Return(nil, errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/status", handler.GetSystemStatus)
			req, _ := http.NewRequest("GET", "/status", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})
}
