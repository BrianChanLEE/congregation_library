package service

import (
	"boock/backGo/internal/models"
	"boock/backGo/internal/repository"
)

// AnnouncementServiceInterface는 공지사항 서비스의 동작을 정의합니다.
type AnnouncementServiceInterface interface {
	CreateAnnouncement(title, content string, authorID int64) error
	GetAllAnnouncements() ([]models.Announcement, error)
	DeleteAnnouncement(id int64) error
	UpdateAnnouncement(id int64, title, content string) error
}

type AnnouncementService struct {
	Repo repository.AnnouncementRepositoryInterface
}

func NewAnnouncementService(repo repository.AnnouncementRepositoryInterface) *AnnouncementService {
	return &AnnouncementService{Repo: repo}
}

func (s *AnnouncementService) CreateAnnouncement(title, content string, authorID int64) error {
	a := &models.Announcement{Title: title, Content: content, AuthorID: authorID}
	return s.Repo.Create(a)
}

func (s *AnnouncementService) GetAllAnnouncements() ([]models.Announcement, error) {
	return s.Repo.GetAll()
}

func (s *AnnouncementService) DeleteAnnouncement(id int64) error {
	return s.Repo.Delete(id)
}

func (s *AnnouncementService) UpdateAnnouncement(id int64, title, content string) error {
	a := &models.Announcement{ID: id, Title: title, Content: content}
	return s.Repo.Update(a)
}
