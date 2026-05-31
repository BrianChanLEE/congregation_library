package service_test

import (
	"boock/backGo/internal/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSystemService_GetSystemErrors 테스트 케이스
func TestSystemService_GetSystemErrors(t *testing.T) {
	mockRepo := new(MockSystemRepository)
	service := service.NewSystemService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockLogs := []map[string]interface{}{{"id": int64(1), "error": "test error"}}
		mockRepo.On("GetErrorLogs").Return(mockLogs, nil).Once()

		logs, err := service.GetSystemErrors()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(logs))
		mockRepo.AssertExpectations(t)
	})

	t.Run("DB 에러 케이스", func(t *testing.T) {
		mockRepo.On("GetErrorLogs").Return([]map[string]interface{}(nil), errors.New("db error")).Once()

		logs, err := service.GetSystemErrors()
		assert.Error(t, err)
		assert.Nil(t, logs)
		assert.Equal(t, "db error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
