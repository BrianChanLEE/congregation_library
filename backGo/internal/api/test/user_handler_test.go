package api_test

import (
	"bytes"
	"boock/backGo/internal/api"
	"boock/backGo/internal/models"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetProfile", func(t *testing.T) {
		mockUser := new(MockUserService)
		handler := api.NewUserHandler(mockUser)

		t.Run("성공", func(t *testing.T) {
			mockUser.On("GetUserProfile", "1").Return(&models.User{ID: 1}, nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/profile", func(gc *gin.Context) {
				gc.Set("userId", int64(1))
				handler.GetProfile(gc)
			})
			req, _ := http.NewRequest("GET", "/profile", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("인증 정보 없음", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/profile", handler.GetProfile)
			req, _ := http.NewRequest("GET", "/profile", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusUnauthorized, w.Code)
		})

		t.Run("사용자 없음", func(t *testing.T) {
			mockUser.On("GetUserProfile", "1").Return(nil, errors.New("not found")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/profile", func(gc *gin.Context) {
				gc.Set("userId", int64(1))
				handler.GetProfile(gc)
			})
			req, _ := http.NewRequest("GET", "/profile", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusNotFound, w.Code)
		})
	})

	t.Run("ChangePassword", func(t *testing.T) {
		mockUser := new(MockUserService)
		handler := api.NewUserHandler(mockUser)

		t.Run("성공", func(t *testing.T) {
			mockUser.On("ChangePassword", int64(1), "old", "new").Return(nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/change-password", func(gc *gin.Context) {
				gc.Set("userId", int64(1))
				handler.ChangePassword(gc)
			})
			input, _ := json.Marshal(gin.H{"current_password": "old", "new_password": "new"})
			req, _ := http.NewRequest("POST", "/change-password", bytes.NewBuffer(input))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("인증 정보 없음", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/change-password", handler.ChangePassword)
			req, _ := http.NewRequest("POST", "/change-password", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusUnauthorized, w.Code)
		})

		t.Run("입력값 오류", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/change-password", func(gc *gin.Context) {
				gc.Set("userId", int64(1))
				handler.ChangePassword(gc)
			})
			req, _ := http.NewRequest("POST", "/change-password", bytes.NewBufferString("invalid"))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("서비스 에러", func(t *testing.T) {
			mockUser.On("ChangePassword", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/change-password", func(gc *gin.Context) {
				gc.Set("userId", int64(1))
				handler.ChangePassword(gc)
			})
			input, _ := json.Marshal(gin.H{"current_password": "old", "new_password": "new"})
			req, _ := http.NewRequest("POST", "/change-password", bytes.NewBuffer(input))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})
}
