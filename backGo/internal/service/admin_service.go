package service

import (
	"boock/backGo/internal/repository"
)

// AdminServiceInterface는 관리자 서비스의 동작을 정의합니다.
type AdminServiceInterface interface {
	GetStats() (int, int, int, error)
}

type AdminService struct {
	Repo repository.AdminRepositoryInterface
}

func NewAdminService(repo repository.AdminRepositoryInterface) *AdminService {
	return &AdminService{Repo: repo}
}

func (s *AdminService) GetStats() (int, int, int, error) {
	return s.Repo.GetStats()
}
