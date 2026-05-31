package service

import (
	"boock/backGo/internal/models"
	"boock/backGo/internal/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

// UserServiceInterface는 사용자 서비스의 동작을 정의합니다.
type UserServiceInterface interface {
	UpdateUserStatus(userID int64, status string) error
	GetPendingUsers() ([]models.User, error)
	GetUserProfile(userIDStr string) (*models.User, error)
	DeleteUser(userID int64) error
	ChangePassword(userID int64, currentPassword, newPassword string) error
}

type UserService struct {
	Repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) UpdateUserStatus(userID int64, status string) error {
	return s.Repo.UpdateStatus(userID, status)
}

func (s *UserService) GetPendingUsers() ([]models.User, error) {
	return s.Repo.GetAllPending()
}

func (s *UserService) GetUserProfile(userIDStr string) (*models.User, error) {
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetByID(userID)
}

func (s *UserService) DeleteUser(userID int64) error {
	return s.Repo.Delete(userID)
}

func (s *UserService) ChangePassword(userID int64, currentPassword, newPassword string) error {
	hashedPassword, err := s.Repo.GetPasswordHash(userID)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(currentPassword)); err != nil {
		return errors.New("현재 비밀번호가 일치하지 않습니다")
	}

	newHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.Repo.UpdatePassword(userID, string(newHash))
}

