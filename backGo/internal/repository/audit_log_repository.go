package repository

import (
	"database/sql"
)

type AuditLogRepository struct {
	db *sql.DB
}

func NewAuditLogRepository(db *sql.DB) *AuditLogRepository {
	return &AuditLogRepository{db: db}
}

func (r *AuditLogRepository) Create(userID *int64, itemID int64, quantity int, actionType, method, details string) error {
	_, err := r.db.Exec("INSERT INTO audit_logs (user_id, item_id, quantity, action_type, method, details) VALUES (?, ?, ?, ?, ?, ?)",
		userID, itemID, quantity, actionType, method, details)
	return err
}
