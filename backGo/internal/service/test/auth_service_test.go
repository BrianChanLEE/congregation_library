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

func TestAuthService_Register(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := service.NewAuthService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything).Return(nil).Once()
		err := svc.Register("test", "password123")
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("실패 케이스 - DB 오류", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything).Return(errors.New("db error")).Once()
		err := svc.Register("test", "password123")
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestAuthService_Login(t *testing.T) {
	mockRepo := new(MockUserRepository)
	svc := service.NewAuthService(mockRepo)
	hashed, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	tests := []struct {
		name          string
		congCode      string
		email         string
		password      string
		mockSetup     func()
		wantErr       bool
	}{
		{
			name:     "성공 케이스",
			congCode: "123",
			email:    "test@jwpub.org",
			password: "password123",
			mockSetup: func() {
				mockRepo.On("GetByJwhubEmailAndCongID", "test@jwpub.org", int64(123)).
					Return(&models.User{JWhubEmail: "test@jwpub.org", PasswordHash: string(hashed), Role: "user", ID: 1}, nil).Once()
			},
			wantErr: false,
		},
		{
			name:     "인증 실패 - 비밀번호 불일치",
			congCode: "123",
			email:    "test@jwpub.org",
			password: "wrongpassword",
			mockSetup: func() {
				mockRepo.On("GetByJwhubEmailAndCongID", "test@jwpub.org", int64(123)).
					Return(&models.User{JWhubEmail: "test@jwpub.org", PasswordHash: string(hashed), Role: "user", ID: 1}, nil).Once()
			},
			wantErr: true,
		},
		{
			name:     "인증 실패 - 사용자 없음",
			congCode: "123",
			email:    "notfound@jwpub.org",
			password: "password123",
			mockSetup: func() {
				mockRepo.On("GetByJwhubEmailAndCongID", "notfound@jwpub.org", int64(123)).
					Return((*models.User)(nil), errors.New("user not found")).Once()
			},
			wantErr: true,
		},
		{
			name:     "로그인 실패 - 잘못된 회중 코드",
			congCode: "abc",
			email:    "test@jwpub.org",
			password: "password123",
			mockSetup: func() {},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			token, err := svc.Login(tt.congCode, tt.email, tt.password)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
