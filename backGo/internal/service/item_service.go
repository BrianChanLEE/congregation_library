package service

import (
	"boock/backGo/internal/models"
	"boock/backGo/internal/repository"
)

// ItemServiceInterface는 품목 서비스의 동작을 정의합니다.
type ItemServiceInterface interface {
	AddItem(name, code string) error
	GetAllItems() ([]models.Item, error)
	GetInventory(congID, itemID int64) (models.Inventory, error)
	DeleteItem(id int64) error
	UpdateItem(id int64, name, code string) error
}

type ItemService struct {
	Repo repository.ItemRepositoryInterface
}

func NewItemService(repo repository.ItemRepositoryInterface) *ItemService {
	return &ItemService{Repo: repo}
}

func (s *ItemService) AddItem(name, code string) error {
	item := &models.Item{Name: name, Code: code}
	return s.Repo.Create(item)
}

func (s *ItemService) GetAllItems() ([]models.Item, error) {
	return s.Repo.GetAll()
}

func (s *ItemService) GetInventory(congID, itemID int64) (models.Inventory, error) {
	return s.Repo.GetInventory(congID, itemID)
}

func (s *ItemService) DeleteItem(id int64) error {
	return s.Repo.Delete(id)
}

func (s *ItemService) UpdateItem(id int64, name, code string) error {
	item := &models.Item{ID: id, Name: name, Code: code}
	return s.Repo.Update(item)
}
