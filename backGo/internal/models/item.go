package models

import "time"

// CatalogItem 카탈로그 조회를 위한 아이템 구조체
type CatalogItem struct {
	ID       int64  `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Category string `json:"category"`
	ImageURL string `json:"imageUrl"`
	Stock    int    `json:"stock"`
}

// Inventory 재고 정보 구조체
type Inventory struct {
	ID    int64  `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

// Item 기본 아이템 정보
type Item struct {
	ID        int64      `json:"id"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	Category  string     `json:"category"`
	FormCode  string     `json:"formCode"`
	NameKo    string     `json:"nameKo"`
	IsSpecial bool       `json:"isSpecial"`
	ImageURL  string     `json:"imageUrl"`
	DeletedAt *time.Time `json:"deletedAt"`
}
