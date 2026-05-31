package models

import "time"

// User 사용자 기본 구조체
type User struct {
	ID             int64      `json:"id"`
	CongregationID int64      `json:"congregationId"`
	Name           string     `json:"name"`
	Role           string     `json:"role"`
	Status         string     `json:"status"`
	Position       string     `json:"position"`
	Phone          string     `json:"phone"`
	Email          string     `json:"email"`
	JWhubEmail     string     `json:"jwhubEmail"`
	PasswordHash   string     `json:"-"`
	CreatedAt      time.Time  `json:"createdAt"`
	DeletedAt      *time.Time `json:"deleted_at"`
}
