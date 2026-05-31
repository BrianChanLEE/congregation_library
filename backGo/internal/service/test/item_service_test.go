package service_test

import (
	"boock/backGo/internal/models"
	"boock/backGo/internal/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestItemService_GetInventory 테스트 케이스
func TestItemService_GetInventory(t *testing.T) {
	mockRepo := new(MockItemRepository)
	svc := service.NewItemService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockInv := models.Inventory{ID: 1, Code: "BOOK01", Name: "테스트 도서", Stock: 10}
		mockRepo.On("GetInventory", int64(1), int64(1)).Return(mockInv, nil).Once()

		inv, err := svc.GetInventory(1, 1)
		assert.NoError(t, err)
		assert.Equal(t, 10, inv.Stock)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DB 에러 케이스", func(t *testing.T) {
		mockRepo.On("GetInventory", int64(1), int64(1)).Return(models.Inventory{}, errors.New("db error")).Once()

		_, err := svc.GetInventory(1, 1)
		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestItemService_AddItem(t *testing.T) {
	mockRepo := new(MockItemRepository)
	svc := service.NewItemService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockRepo.On("Create", &models.Item{Name: "테스트", Code: "CODE01"}).Return(nil).Once()
		err := svc.AddItem("테스트", "CODE01")
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestItemService_GetAllItems(t *testing.T) {
	mockRepo := new(MockItemRepository)
	svc := service.NewItemService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockItems := []models.Item{{ID: 1, Name: "테스트", Code: "CODE01"}}
		mockRepo.On("GetAll").Return(mockItems, nil).Once()
		items, err := svc.GetAllItems()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(items))
		mockRepo.AssertExpectations(t)
	})
}

func TestItemService_DeleteItem(t *testing.T) {
	mockRepo := new(MockItemRepository)
	svc := service.NewItemService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockRepo.On("Delete", int64(1)).Return(nil).Once()
		err := svc.DeleteItem(1)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestItemService_UpdateItem(t *testing.T) {
	mockRepo := new(MockItemRepository)
	svc := service.NewItemService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockItem := &models.Item{ID: 1, Name: "테스트", Code: "CODE01"}
		mockRepo.On("Update", mockItem).Return(nil).Once()
		err := svc.UpdateItem(1, "테스트", "CODE01")
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}
