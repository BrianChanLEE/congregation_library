package service

import (
	"boock/backGo/internal/models"
	"boock/backGo/internal/repository"
)

// CatalogServiceInterface는 카탈로그 서비스의 동작을 정의합니다.
type CatalogServiceInterface interface {
	GetCatalog() ([]models.CatalogItem, error)
	GetCategories() ([]string, error)
}

// CatalogService 카탈로그 비즈니스 로직을 담당하는 서비스
type CatalogService struct {
	itemRepo repository.ItemRepositoryInterface
}

// NewCatalogService 새로운 CatalogService 생성
func NewCatalogService(repo repository.ItemRepositoryInterface) *CatalogService {
	return &CatalogService{
		itemRepo: repo,
	}
}

// GetCatalog 전체 출판물 목록 반환
func (s *CatalogService) GetCatalog() ([]models.CatalogItem, error) {
	items, err := s.itemRepo.GetCatalog()
	if err != nil {
		return nil, err
	}

	var catalog []models.CatalogItem
	for _, item := range items {
		catalog = append(catalog, models.CatalogItem{
			ID:   item.ID,
			Code: item.Code,
			Name: item.Name,
		})
	}
	return catalog, nil
}

func (s *CatalogService) GetCategories() ([]string, error) {
	return s.itemRepo.GetCategories()
}
