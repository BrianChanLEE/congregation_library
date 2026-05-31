package service_test

import (
	"boock/backGo/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockActivityLogRepository struct {
	mock.Mock
}

func (m *MockActivityLogRepository) Create(log *models.ActivityLog) error {
	args := m.Called(log)
	return args.Error(0)
}

func (m *MockActivityLogRepository) GetAll() ([]models.ActivityLog, error) {
	args := m.Called()
	return args.Get(0).([]models.ActivityLog), args.Error(1)
}

func (m *MockActivityLogRepository) UpdateType(id int64, logType string) error {
	args := m.Called(id, logType)
	return args.Error(0)
}

func (m *MockActivityLogRepository) GetDetailed() ([]map[string]interface{}, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]map[string]interface{}), args.Error(1)
}

type MockAnnouncementRepository struct {
	mock.Mock
}

func (m *MockAnnouncementRepository) Create(a *models.Announcement) error {
	args := m.Called(a)
	return args.Error(0)
}

func (m *MockAnnouncementRepository) GetAll() ([]models.Announcement, error) {
	args := m.Called()
	return args.Get(0).([]models.Announcement), args.Error(1)
}

func (m *MockAnnouncementRepository) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockAnnouncementRepository) Update(a *models.Announcement) error {
	args := m.Called(a)
	return args.Error(0)
}

type MockItemRepository struct {
	mock.Mock
}

func (m *MockItemRepository) Create(item *models.Item) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockItemRepository) GetAll() ([]models.Item, error) {
	args := m.Called()
	return args.Get(0).([]models.Item), args.Error(1)
}

func (m *MockItemRepository) GetCatalog() ([]models.Item, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Item), args.Error(1)
}

func (m *MockItemRepository) GetInventory(congID, itemID int64) (models.Inventory, error) {
	args := m.Called(congID, itemID)
	if args.Get(0) == nil {
		return models.Inventory{}, args.Error(1)
	}
	return args.Get(0).(models.Inventory), args.Error(1)
}

func (m *MockItemRepository) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockItemRepository) Update(item *models.Item) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockItemRepository) GetCategories() ([]string, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]string), args.Error(1)
}

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(id int64) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetByJwhubEmailAndCongID(email string, congID int64) (*models.User, error) {
	args := m.Called(email, congID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateStatus(userID int64, status string) error {
	args := m.Called(userID, status)
	return args.Error(0)
}

func (m *MockUserRepository) GetAllPending() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) Delete(userID int64) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *MockUserRepository) GetPasswordHash(userID int64) (string, error) {
	args := m.Called(userID)
	return args.String(0), args.Error(1)
}

func (m *MockUserRepository) UpdatePassword(userID int64, hashedNewPassword string) error {
	args := m.Called(userID, hashedNewPassword)
	return args.Error(0)
}

type MockSystemRepository struct {
	mock.Mock
}

func (m *MockSystemRepository) GetErrorLogs() ([]map[string]interface{}, error) {
	args := m.Called()
	return args.Get(0).([]map[string]interface{}), args.Error(1)
}
