package service_test

import (
	"boock/backGo/internal/service"
	"boock/backGo/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestActivityLogService_CreateLog 테스트 케이스
func TestActivityLogService_CreateLog(t *testing.T) {
	mockRepo := new(MockActivityLogRepository)
	service := service.NewActivityLogService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything).Return(nil)

		err := service.CreateLog(1, 1, 10, "IN", "WEB", "테스트")
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

// TestActivityLogService_GetAllLogs 테스트 케이스
func TestActivityLogService_GetAllLogs(t *testing.T) {
	mockRepo := new(MockActivityLogRepository)
	service := service.NewActivityLogService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockLogs := []models.ActivityLog{{ID: 1, UserID: 1, ItemID: 1, Quantity: 10}}
		mockRepo.On("GetAll").Return(mockLogs, nil)

		logs, err := service.GetAllLogs()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(logs))
		mockRepo.AssertExpectations(t)
	})
}

func TestActivityLogService_CancelLog(t *testing.T) {
	mockRepo := new(MockActivityLogRepository)
	svc := service.NewActivityLogService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockRepo.On("UpdateType", int64(1), "CANCEL").Return(nil).Once()
		err := svc.CancelLog(1)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestActivityLogService_GetDetailedLogs(t *testing.T) {
	mockRepo := new(MockActivityLogRepository)
	svc := service.NewActivityLogService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockDetailed := []map[string]interface{}{{"id": 1}}
		mockRepo.On("GetDetailed").Return(mockDetailed, nil).Once()
		logs, err := svc.GetDetailedLogs()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(logs))
		mockRepo.AssertExpectations(t)
	})
}
