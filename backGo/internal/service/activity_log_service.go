package service

import (
	"boock/backGo/internal/models"
	"boock/backGo/internal/repository"
)

// ActivityLogServiceInterface는 활동 로그 서비스의 동작을 정의합니다.
type ActivityLogServiceInterface interface {
	CreateLog(userID, itemID int64, quantity int, logType, method, memo string) error
	GetAllLogs() ([]models.ActivityLog, error)
	CancelLog(id int64) error
	GetDetailedLogs() ([]map[string]interface{}, error)
}

type ActivityLogService struct {
	Repo repository.ActivityLogRepositoryInterface
}

func NewActivityLogService(repo repository.ActivityLogRepositoryInterface) *ActivityLogService {
	return &ActivityLogService{Repo: repo}
}

func (s *ActivityLogService) CreateLog(userID, itemID int64, quantity int, logType, method, memo string) error {
	log := &models.ActivityLog{
		UserID:   userID,
		ItemID:   itemID,
		Quantity: quantity,
		Type:     logType,
		Method:   method,
		Memo:     memo,
	}
	return s.Repo.Create(log)
}

func (s *ActivityLogService) GetAllLogs() ([]models.ActivityLog, error) {
	return s.Repo.GetAll()
}

func (s *ActivityLogService) CancelLog(id int64) error {
	return s.Repo.UpdateType(id, "CANCEL")
}

func (s *ActivityLogService) GetDetailedLogs() ([]map[string]interface{}, error) {
	return s.Repo.GetDetailed()
}
