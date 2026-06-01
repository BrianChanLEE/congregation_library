package service_test

import (
	"boock/backGo/internal/models"
	"boock/backGo/internal/service"

	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCatalogService_GetCatalog 테스트 케이스
func TestCatalogService_GetCatalog(t *testing.T) {
	mockRepo := new(MockItemRepository)
	svc := service.NewCatalogService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockItems := []models.Item{
			{ID: 1, Code: "BOOK01", Name: "테스트 도서"},
		}
		mockRepo.On("GetCatalog").Return(mockItems, nil).Once()

		catalog, err := svc.GetCatalog()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(catalog))
		assert.Equal(t, "테스트 도서", catalog[0].Name)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DB 에러 케이스", func(t *testing.T) {
		mockRepo.On("GetCatalog").Return(nil, errors.New("db error")).Once()

		catalog, err := svc.GetCatalog()
		assert.Error(t, err)
		assert.Nil(t, catalog)
		assert.Equal(t, "db error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestCatalogService_GetCategories(t *testing.T) {
	mockRepo := new(MockItemRepository)
	svc := service.NewCatalogService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockCats := []string{"Books", "Magazines"}
		mockRepo.On("GetCategories").Return(mockCats, nil).Once()

		cats, err := svc.GetCategories()
		assert.NoError(t, err)
		assert.Equal(t, 2, len(cats))
		mockRepo.AssertExpectations(t)
	})

	t.Run("에러 케이스", func(t *testing.T) {
		mockRepo.On("GetCategories").Return(nil, errors.New("error")).Once()

		cats, err := svc.GetCategories()
		assert.Error(t, err)
		assert.Nil(t, cats)
		mockRepo.AssertExpectations(t)
	})
}
