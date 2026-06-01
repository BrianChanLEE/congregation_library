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

type MockAnnouncementService struct {
	mock.Mock
}

func (m *MockAnnouncementService) CreateAnnouncement(title, content string, authorID int64) error {
	args := m.Called(title, content, authorID)
	return args.Error(0)
}

func (m *MockAnnouncementService) GetAllAnnouncements() ([]models.Announcement, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]models.Announcement), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAnnouncementService) DeleteAnnouncement(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockAnnouncementService) UpdateAnnouncement(id int64, title, content string) error {
	args := m.Called(id, title, content)
	return args.Error(0)
}

func TestAnnouncementHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetAnnouncements", func(t *testing.T) {
		mockSvc := new(MockAnnouncementService)
		handler := api.NewAnnouncementHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("GetAllAnnouncements").Return([]models.Announcement{{ID: 1}}, nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/anns", handler.GetAnnouncements)
			req, _ := http.NewRequest("GET", "/anns", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("에러", func(t *testing.T) {
			mockSvc.On("GetAllAnnouncements").Return(nil, errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/anns", handler.GetAnnouncements)
			req, _ := http.NewRequest("GET", "/anns", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})

	t.Run("AddAnnouncement", func(t *testing.T) {
		mockSvc := new(MockAnnouncementService)
		handler := api.NewAnnouncementHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("CreateAnnouncement", "T", "C", int64(1)).Return(nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/anns", handler.AddAnnouncement)
			input, _ := json.Marshal(gin.H{"title": "T", "content": "C", "author_id": 1})
			req, _ := http.NewRequest("POST", "/anns", bytes.NewBuffer(input))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusCreated, w.Code)
		})

		t.Run("입력값 오류", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/anns", handler.AddAnnouncement)
			req, _ := http.NewRequest("POST", "/anns", bytes.NewBufferString("invalid"))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("서비스 에러", func(t *testing.T) {
			mockSvc.On("CreateAnnouncement", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/anns", handler.AddAnnouncement)
			input, _ := json.Marshal(gin.H{"title": "T"})
			req, _ := http.NewRequest("POST", "/anns", bytes.NewBuffer(input))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})

	t.Run("UpdateAnnouncement", func(t *testing.T) {
		mockSvc := new(MockAnnouncementService)
		handler := api.NewAnnouncementHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("UpdateAnnouncement", int64(1), "T", "C").Return(nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.PUT("/anns/:id", handler.UpdateAnnouncement)
			input, _ := json.Marshal(gin.H{"title": "T", "content": "C"})
			req, _ := http.NewRequest("PUT", "/anns/1", bytes.NewBuffer(input))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("잘못된 ID", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.PUT("/anns/:id", handler.UpdateAnnouncement)
			req, _ := http.NewRequest("PUT", "/anns/abc", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("입력값 오류", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.PUT("/anns/:id", handler.UpdateAnnouncement)
			req, _ := http.NewRequest("PUT", "/anns/1", bytes.NewBufferString("invalid"))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("서비스 에러", func(t *testing.T) {
			mockSvc.On("UpdateAnnouncement", int64(1), mock.Anything, mock.Anything).Return(errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.PUT("/anns/:id", handler.UpdateAnnouncement)
			input, _ := json.Marshal(gin.H{"title": "T"})
			req, _ := http.NewRequest("PUT", "/anns/1", bytes.NewBuffer(input))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})

	t.Run("DeleteAnnouncement", func(t *testing.T) {
		mockSvc := new(MockAnnouncementService)
		handler := api.NewAnnouncementHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("DeleteAnnouncement", int64(1)).Return(nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.DELETE("/anns/:id", handler.DeleteAnnouncement)
			req, _ := http.NewRequest("DELETE", "/anns/1", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("잘못된 ID", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.DELETE("/anns/:id", handler.DeleteAnnouncement)
			req, _ := http.NewRequest("DELETE", "/anns/abc", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("서비스 에러", func(t *testing.T) {
			mockSvc.On("DeleteAnnouncement", int64(1)).Return(errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.DELETE("/anns/:id", handler.DeleteAnnouncement)
			req, _ := http.NewRequest("DELETE", "/anns/1", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})
}
