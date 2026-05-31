package api_test

import (
	"boock/backGo/internal/api"
	"boock/backGo/internal/models"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCatalogService struct {
	mock.Mock
}

func (m *MockCatalogService) GetCatalog() ([]models.CatalogItem, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]models.CatalogItem), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCatalogService) GetCategories() ([]string, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]string), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestCatalogHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GetCatalog", func(t *testing.T) {
		mockSvc := new(MockCatalogService)
		handler := api.NewCatalogHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("GetCatalog").Return([]models.CatalogItem{{ID: 1}}, nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/catalog", handler.GetCatalog)
			req, _ := http.NewRequest("GET", "/catalog", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("에러", func(t *testing.T) {
			mockSvc.On("GetCatalog").Return(nil, errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/catalog", handler.GetCatalog)
			req, _ := http.NewRequest("GET", "/catalog", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})

	t.Run("GetCategories", func(t *testing.T) {
		mockSvc := new(MockCatalogService)
		handler := api.NewCatalogHandler(mockSvc)

		t.Run("성공", func(t *testing.T) {
			mockSvc.On("GetCategories").Return([]string{"A", "B"}, nil).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/categories", handler.GetCategories)
			req, _ := http.NewRequest("GET", "/categories", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("에러", func(t *testing.T) {
			mockSvc.On("GetCategories").Return(nil, errors.New("error")).Once()
			w := httptest.NewRecorder()
			_, r := gin.CreateTestContext(w)
			r.GET("/categories", handler.GetCategories)
			req, _ := http.NewRequest("GET", "/categories", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	})
}
