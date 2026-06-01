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

type MockActivityLogService struct {
	mock.Mock
}

func (m *MockActivityLogService) CreateLog(userID, itemID int64, quantity int, logType, method, memo string) error {
	args := m.Called(userID, itemID, quantity, logType, method, memo)
	return args.Error(0)
}

func (m *MockActivityLogService) GetAllLogs() ([]models.ActivityLog, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]models.ActivityLog), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockActivityLogService) CancelLog(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockActivityLogService) GetDetailedLogs() ([]map[string]interface{}, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]map[string]interface{}), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestTransactionHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("AddTransaction", func(t *testing.T) {
		mockSvc := new(MockActivityLogService)
		handler := api.NewTransactionHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("CreateLog", int64(1), int64(10), 5, "OUT", "GIFT", "Test memo").Return(nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/transactions", handler.AddTransaction)
			input, _ := json.Marshal(gin.H{"user_id": 1, "item_id": 10, "quantity": 5, "type": "OUT", "method": "GIFT", "memo": "Test memo"})
			req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(input))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusCreated, w.Code)
		})

		t.Run("입력 오류", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/transactions", handler.AddTransaction)
			req, _ := http.NewRequest("POST", "/transactions", bytes.NewBufferString("invalid"))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("서비스 에러", func(t *testing.T) {
			mockSvc.On("CreateLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/transactions", handler.AddTransaction)
			input, _ := json.Marshal(gin.H{"user_id": 1})
			req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(input))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})

	t.Run("CancelTransaction", func(t *testing.T) {
		mockSvc := new(MockActivityLogService)
		handler := api.NewTransactionHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("CancelLog", int64(1)).Return(nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/cancel/:id", handler.CancelTransaction)
			req, _ := http.NewRequest("POST", "/cancel/1", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("잘못된 ID", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/cancel/:id", handler.CancelTransaction)
			req, _ := http.NewRequest("POST", "/cancel/abc", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("서비스 에러", func(t *testing.T) {
			mockSvc.On("CancelLog", int64(1)).Return(errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/cancel/:id", handler.CancelTransaction)
			req, _ := http.NewRequest("POST", "/cancel/1", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})

	t.Run("GetAuditLogs", func(t *testing.T) {
		mockSvc := new(MockActivityLogService)
		handler := api.NewTransactionHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("GetAllLogs").Return([]models.ActivityLog{{ID: 1}}, nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/history", handler.GetAuditLogs)
			req, _ := http.NewRequest("GET", "/history", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("에러", func(t *testing.T) {
			mockSvc.On("GetAllLogs").Return(nil, errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/history", handler.GetAuditLogs)
			req, _ := http.NewRequest("GET", "/history", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})

	t.Run("GetOldAuditLogs", func(t *testing.T) {
		mockSvc := new(MockActivityLogService)
		handler := api.NewTransactionHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("GetDetailedLogs").Return([]map[string]interface{}{{"id": 1}}, nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/old-history", handler.GetOldAuditLogs)
			req, _ := http.NewRequest("GET", "/old-history", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("에러", func(t *testing.T) {
			mockSvc.On("GetDetailedLogs").Return(nil, errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/old-history", handler.GetOldAuditLogs)
			req, _ := http.NewRequest("GET", "/old-history", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})
}
