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

type MockItemService struct {
	mock.Mock
}

func (m *MockItemService) AddItem(name, code string) error {
	args := m.Called(name, code)
	return args.Error(0)
}

func (m *MockItemService) GetAllItems() ([]models.Item, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]models.Item), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockItemService) GetInventory(congID, itemID int64) (models.Inventory, error) {
	args := m.Called(congID, itemID)
	return args.Get(0).(models.Inventory), args.Error(1)
}

func (m *MockItemService) DeleteItem(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockItemService) UpdateItem(id int64, name, code string) error {
	args := m.Called(id, name, code)
	return args.Error(0)
}

func TestItemHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("AddItem", func(t *testing.T) {
		mockSvc := new(MockItemService)
		handler := api.NewItemHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("AddItem", "Item1", "C1").Return(nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/items", handler.AddItem)
			input, _ := json.Marshal(models.Item{Name: "Item1", Code: "C1"})
			req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(input))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusCreated, w.Code)
		})

		t.Run("입력 오류", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/items", handler.AddItem)
			req, _ := http.NewRequest("POST", "/items", bytes.NewBufferString("invalid"))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("서비스 에러", func(t *testing.T) {
			mockSvc.On("AddItem", mock.Anything, mock.Anything).Return(errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.POST("/items", handler.AddItem)
			input, _ := json.Marshal(models.Item{Name: "Item1"})
			req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(input))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})

	t.Run("GetItems", func(t *testing.T) {
		mockSvc := new(MockItemService)
		handler := api.NewItemHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("GetAllItems").Return([]models.Item{{ID: 1}}, nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/items", handler.GetItems)
			req, _ := http.NewRequest("GET", "/items", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("에러", func(t *testing.T) {
			mockSvc.On("GetAllItems").Return(nil, errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/items", handler.GetItems)
			req, _ := http.NewRequest("GET", "/items", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})

	t.Run("UpdateItem", func(t *testing.T) {
		mockSvc := new(MockItemService)
		handler := api.NewItemHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("UpdateItem", int64(1), "N", "C").Return(nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.PUT("/items/:id", handler.UpdateItem)
			input, _ := json.Marshal(gin.H{"name": "N", "code": "C"})
			req, _ := http.NewRequest("PUT", "/items/1", bytes.NewBuffer(input))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("잘못된 ID", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.PUT("/items/:id", handler.UpdateItem)
			req, _ := http.NewRequest("PUT", "/items/abc", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
		
		t.Run("바인딩 에러", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.PUT("/items/:id", handler.UpdateItem)
			req, _ := http.NewRequest("PUT", "/items/1", bytes.NewBufferString("invalid"))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("서비스 에러", func(t *testing.T) {
			mockSvc.On("UpdateItem", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.PUT("/items/:id", handler.UpdateItem)
			input, _ := json.Marshal(gin.H{"name": "N"})
			req, _ := http.NewRequest("PUT", "/items/1", bytes.NewBuffer(input))
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})

	t.Run("DeleteItem", func(t *testing.T) {
		mockSvc := new(MockItemService)
		handler := api.NewItemHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("DeleteItem", int64(1)).Return(nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.DELETE("/items/:id", handler.DeleteItem)
			req, _ := http.NewRequest("DELETE", "/items/1", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("잘못된 ID", func(t *testing.T) {
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.DELETE("/items/:id", handler.DeleteItem)
			req, _ := http.NewRequest("DELETE", "/items/abc", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusBadRequest, w.Code)
		})

		t.Run("서비스 에러", func(t *testing.T) {
			mockSvc.On("DeleteItem", int64(1)).Return(errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.DELETE("/items/:id", handler.DeleteItem)
			req, _ := http.NewRequest("DELETE", "/items/1", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})
}

func TestInventoryHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetInventory", func(t *testing.T) {
		mockSvc := new(MockItemService)
		handler := api.NewItemHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("GetInventory", int64(1), int64(10)).Return(models.Inventory{Stock: 100}, nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/inventory", handler.GetInventory)
			req, _ := http.NewRequest("GET", "/inventory?cong_id=1&item_id=10", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("에러", func(t *testing.T) {
			mockSvc.On("GetInventory", mock.Anything, mock.Anything).Return(models.Inventory{}, errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/inventory", handler.GetInventory)
			req, _ := http.NewRequest("GET", "/inventory?cong_id=1&item_id=10", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})
}
