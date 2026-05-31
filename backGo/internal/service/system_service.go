package service

import (
	"boock/backGo/internal/repository"
)

// SystemServiceInterface는 시스템 서비스의 동작을 정의합니다.
type SystemServiceInterface interface {
	GetSystemErrors() ([]map[string]interface{}, error)
}

type SystemService struct {
	Repo repository.SystemRepositoryInterface
}

func NewSystemService(repo repository.SystemRepositoryInterface) *SystemService {
	return &SystemService{Repo: repo}
}

func (s *SystemService) GetSystemErrors() ([]map[string]interface{}, error) {
	return s.Repo.GetErrorLogs()
}
