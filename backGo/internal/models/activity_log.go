package models

import "time"

// ActivityLog 활동 로그 구조체
type ActivityLog struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	ItemID    int64     `json:"item_id"`
	Quantity  int       `json:"quantity"`
	Type      string    `json:"type"`
	Method    string    `json:"method"`
	Memo      string    `json:"memo"`
	CreatedAt time.Time `json:"created_at"`
	// Join 필드
	UserName string `json:"user_name"`
	ItemName string `json:"item_name"`
}
