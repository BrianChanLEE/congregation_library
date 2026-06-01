package service_test

import (
	"boock/backGo/internal/models"
	"boock/backGo/internal/service"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestAnnouncementService_CreateAnnouncement 테스트 케이스
func TestAnnouncementService_CreateAnnouncement(t *testing.T) {
	mockRepo := new(MockAnnouncementRepository)
	svc := service.NewAnnouncementService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything).Return(nil)

		err := svc.CreateAnnouncement("테스트", "내용", 1)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

// TestAnnouncementService_GetAllAnnouncements 테스트 케이스
func TestAnnouncementService_GetAllAnnouncements(t *testing.T) {
	mockRepo := new(MockAnnouncementRepository)
	svc := service.NewAnnouncementService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockRepo.On("GetAll").Return([]models.Announcement{{ID: 1, Title: "공지"}}, nil)

		anns, err := svc.GetAllAnnouncements()
		assert.NoError(t, err)
		assert.Equal(t, 1, len(anns))
		assert.Equal(t, "공지", anns[0].Title)
		mockRepo.AssertExpectations(t)
	})
}

// ...
// TestAnnouncementService_DeleteAnnouncement 테스트 케이스
func TestAnnouncementService_DeleteAnnouncement(t *testing.T) {
	mockRepo := new(MockAnnouncementRepository)
	svc := service.NewAnnouncementService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		mockRepo.On("Delete", int64(1)).Return(nil)

		err := svc.DeleteAnnouncement(1)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

// TestAnnouncementService_UpdateAnnouncement 테스트 케이스
func TestAnnouncementService_UpdateAnnouncement(t *testing.T) {
	mockRepo := new(MockAnnouncementRepository)
	svc := service.NewAnnouncementService(mockRepo)

	t.Run("성공 케이스", func(t *testing.T) {
		a := &models.Announcement{ID: 1, Title: "수정", Content: "수정 내용"}
		mockRepo.On("Update", a).Return(nil)

		err := svc.UpdateAnnouncement(1, "수정", "수정 내용")
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}
