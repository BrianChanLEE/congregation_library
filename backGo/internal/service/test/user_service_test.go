package service_test

import (
	"boock/backGo/internal/models"
	"boock/backGo/internal/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)
// ...

// TestUserService_UpdateUserStatus 테스트 케이스
func TestUserService_UpdateUserStatus(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := service.NewUserService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockRepo.On("UpdateStatus", int64(1), "APPROVED").Return(nil).Once()

		err := svc.UpdateUserStatus(1, "APPROVED")
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DB 에러 케이스", func(t *testing.T) {
		mockRepo.On("UpdateStatus", int64(1), "APPROVED").Return(errors.New("db error")).Once()

		err := svc.UpdateUserStatus(1, "APPROVED")
		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

// TestUserService_GetPendingUsers 테스트 케이스
func TestUserService_GetPendingUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := service.NewUserService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockUsers := []models.User{{ID: 1, Name: "Test User", Status: "PENDING"}}
		mockRepo.On("GetAllPending").Return(mockUsers, nil)

		users, err := svc.GetPendingUsers()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(users))
		mockRepo.AssertExpectations(t)
	})
}

// TestUserService_GetUserProfile 테스트 케이스
func TestUserService_GetUserProfile(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := service.NewUserService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockUser := &models.User{ID: 1, Name: "Test User"}
		mockRepo.On("GetByID", int64(1)).Return(mockUser, nil)

		user, err := svc.GetUserProfile("1")
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "Test User", user.Name)
		mockRepo.AssertExpectations(t)
	})

	t.Run("잘못된 ID 케이스", func(t *testing.T) {
		user, err := svc.GetUserProfile("invalid")
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}

// TestUserService_ChangePassword 테스트 케이스
func TestUserService_ChangePassword(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := service.NewUserService(mockRepo)
	hashed, _ := bcrypt.GenerateFromPassword([]byte("oldpassword"), bcrypt.DefaultCost)

	t.Run("성공 케이스", func(t *testing.T) {
		mockRepo.On("GetPasswordHash", int64(1)).Return(string(hashed), nil).Once()
		mockRepo.On("UpdatePassword", int64(1), mock.AnythingOfType("string")).Return(nil).Once()

		err := svc.ChangePassword(1, "oldpassword", "newpassword")
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
// ... (이전 코드)
	t.Run("비밀번호 불일치 케이스", func(t *testing.T) {
		mockRepo.On("GetPasswordHash", int64(1)).Return(string(hashed), nil).Once()

		err := svc.ChangePassword(1, "wrongpassword", "newpassword")
		assert.Error(t, err)
		assert.Equal(t, "현재 비밀번호가 일치하지 않습니다", err.Error())
		mockRepo.AssertExpectations(t)
	})
}


func TestUserService_DeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := service.NewUserService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockRepo.On("Delete", int64(1)).Return(nil).Once()
		err := svc.DeleteUser(1)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

// Note: bcrypt error path unreachable via normal input. Need simulate.
// bcrypt.GenerateFromPassword fail only on extreme cost or long input.
// Mocking bcrypt hard. Add test for GetPasswordHash DB error.

func TestUserService_ChangePassword_RepoError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := service.NewUserService(mockRepo)

	t.Run("GetPasswordHash DB 에러 케이스", func(t *testing.T) {
		mockRepo.On("GetPasswordHash", int64(1)).Return("", errors.New("db error")).Once()

		err := svc.ChangePassword(1, "oldpassword", "newpassword")
		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdatePassword DB 에러 케이스", func(t *testing.T) {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("oldpassword"), bcrypt.DefaultCost)
		mockRepo.On("GetPasswordHash", int64(1)).Return(string(hashed), nil).Once()
		mockRepo.On("UpdatePassword", int64(1), mock.AnythingOfType("string")).Return(errors.New("db error")).Once()

		err := svc.ChangePassword(1, "oldpassword", "newpassword")
		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_DeleteUser_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := service.NewUserService(mockRepo)

	t.Run("DB 에러 케이스", func(t *testing.T) {
		mockRepo.On("Delete", int64(1)).Return(errors.New("db error")).Once()
		err := svc.DeleteUser(1)
		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
