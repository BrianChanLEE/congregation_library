package models

import "time"

// Transaction 트랜잭션 기록 구조체
type Transaction struct {
	ID         int64     `json:"id"`
	FromCongID *int64    `json:"fromCongId"`
	ToCongID   int64     `json:"toCongId"`
	ItemID     int64     `json:"itemId"`
	Quantity   int       `json:"quantity"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"createdAt"`
}

// TransactionRequest 거래 생성 요청
type TransactionRequest struct {
	FromCongID *int64 `json:"from_cong_id"`
	ToCongID   int64  `json:"to_cong_id" binding:"required"`
	ItemID     int64  `json:"item_id" binding:"required"`
	Quantity   int    `json:"quantity" binding:"required"`
	Type       string `json:"type" binding:"required"`
}
