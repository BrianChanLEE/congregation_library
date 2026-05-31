package service_test

import (
	"boock/backGo/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAdminRepository struct {
	mock.Mock
}

func (m *MockAdminRepository) GetStats() (int, int, int, error) {
	args := m.Called()
	return args.Int(0), args.Int(1), args.Int(2), args.Error(3)
}

func TestAdminService_GetStats(t *testing.T) {
	mockRepo := new(MockAdminRepository)
	service := service.NewAdminService(mockRepo)

	t.Run("성공", func(t *testing.T) {
		mockRepo.On("GetStats").Return(10, 5, 2, nil)
		items, act, users, err := service.GetStats()
		assert.NoError(t, err)
		assert.Equal(t, 10, items)
		assert.Equal(t, 5, act)
		assert.Equal(t, 2, users)
	})
}
