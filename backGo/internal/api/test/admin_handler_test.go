package api_test

import (
	"boock/backGo/internal/api"
	"boock/backGo/internal/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAdminService struct {
	mock.Mock
}

func (m *MockAdminService) GetStats() (int, int, int, error) {
	args := m.Called()
	return args.Int(0), args.Int(1), args.Int(2), args.Error(3)
}

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) UpdateUserStatus(userID int64, status string) error {
	args := m.Called(userID, status)
	return args.Error(0)
}

func (m *MockUserService) GetPendingUsers() ([]models.User, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserService) GetUserProfile(userIDStr string) (*models.User, error) {
	args := m.Called(userIDStr)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserService) DeleteUser(userID int64) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *MockUserService) ChangePassword(userID int64, cur, new string) error {
	args := m.Called(userID, cur, new)
	return args.Error(0)
}

func TestAdminHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetAdminStats", func(t *testing.T) {
		mockAdmin := new(MockAdminService)
		mockUser := new(MockUserService)
		handler := api.NewAdminHandler(mockAdmin, mockUser)

		t.Run("성공", func(t *testing.T) {
			mockAdmin.On("GetStats").Return(10, 5, 2, nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/stats", handler.GetAdminStats)
			req, _ := http.NewRequest("GET", "/stats", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("에러", func(t *testing.T) {
			mockAdmin.On("GetStats").Return(0, 0, 0, errors.New("db error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/stats", handler.GetAdminStats)
			req, _ := http.NewRequest("GET", "/stats", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})

	t.Run("GetPendingUsers", func(t *testing.T) {
		mockAdmin := new(MockAdminService)
		mockUser := new(MockUserService)
		handler := api.NewAdminHandler(mockAdmin, mockUser)

		t.Run("성공", func(t *testing.T) {
			mockUser.On("GetPendingUsers").Return([]models.User{{ID: 1}}, nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/pending", handler.GetPendingUsers)
			req, _ := http.NewRequest("GET", "/pending", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("에러", func(t *testing.T) {
			mockUser.On("GetPendingUsers").Return(nil, errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/pending", handler.GetPendingUsers)
			req, _ := http.NewRequest("GET", "/pending", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})

	t.Run("UpdateUserStatus", func(t *testing.T) {
		mockAdmin := new(MockAdminService)
		mockUser := new(MockUserService)
		handler := api.NewAdminHandler(mockAdmin, mockUser)

		t.Run("성공", func(t *testing.T) {
			mockUser.On("UpdateUserStatus", int64(1), "APPROVED").Return(nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.PUT("/users/:id", handler.UpdateUserStatus)
			input, _ := json.Marshal(gin.H{"status": "APPROVED"})
			req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(input))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("잘못된 ID", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.PUT("/users/:id", handler.UpdateUserStatus)
			req, _ := http.NewRequest("PUT", "/users/abc", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("바인딩 에러", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.PUT("/users/:id", handler.UpdateUserStatus)
			req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBufferString("invalid"))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("서비스 에러", func(t *testing.T) {
			mockUser.On("UpdateUserStatus", int64(1), "APPROVED").Return(errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.PUT("/users/:id", handler.UpdateUserStatus)
			input, _ := json.Marshal(gin.H{"status": "APPROVED"})
			req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(input))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})

	t.Run("DeleteUser", func(t *testing.T) {
		mockAdmin := new(MockAdminService)
		mockUser := new(MockUserService)
		handler := api.NewAdminHandler(mockAdmin, mockUser)

		t.Run("성공", func(t *testing.T) {
			mockUser.On("DeleteUser", int64(1)).Return(nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.DELETE("/users/:id", handler.DeleteUser)
			req, _ := http.NewRequest("DELETE", "/users/1", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("에러", func(t *testing.T) {
			mockUser.On("DeleteUser", int64(1)).Return(errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.DELETE("/users/:id", handler.DeleteUser)
			req, _ := http.NewRequest("DELETE", "/users/1", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})

		t.Run("잘못된 ID", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.DELETE("/users/:id", handler.DeleteUser)
			req, _ := http.NewRequest("DELETE", "/users/abc", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	})

	t.Run("GetAlerts", func(t *testing.T) {
		handler := api.NewAdminHandler(nil, nil)
		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)
		r.GET("/alerts", handler.GetAlerts)
		req, _ := http.NewRequest("GET", "/alerts", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
