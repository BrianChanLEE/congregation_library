package repository

import (
	"boock/backGo/internal/db"
	"boock/backGo/internal/models"
)

type ActivityLogRepositoryInterface interface {
	Create(log *models.ActivityLog) error
	GetAll() ([]models.ActivityLog, error)
	UpdateType(id int64, logType string) error
	GetDetailed() ([]map[string]interface{}, error)
}

type ActivityLogRepository struct{}

func (r *ActivityLogRepository) Create(log *models.ActivityLog) error {
	query := "INSERT INTO activity_logs (user_id, item_id, quantity, type, method, memo) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := db.DB.Exec(query, log.UserID, log.ItemID, log.Quantity, log.Type, log.Method, log.Memo)
	return err
}

func (r *ActivityLogRepository) GetAll() ([]models.ActivityLog, error) {
	query := "SELECT al.id, al.user_id, al.item_id, al.quantity, al.type, al.method, al.memo, al.created_at, u.name, i.name FROM activity_logs al JOIN users u ON al.user_id = u.id JOIN items i ON al.item_id = i.id ORDER BY al.created_at DESC"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []models.ActivityLog
	for rows.Next() {
		var log models.ActivityLog
		if err := rows.Scan(&log.ID, &log.UserID, &log.ItemID, &log.Quantity, &log.Type, &log.Method, &log.Memo, &log.CreatedAt, &log.UserName, &log.ItemName); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}

func (r *ActivityLogRepository) UpdateType(id int64, logType string) error {
	_, err := db.DB.Exec("UPDATE activity_logs SET type = ? WHERE id = ?", logType, id)
	return err
}

func (r *ActivityLogRepository) GetDetailed() ([]map[string]interface{}, error) {
	// This is a simplified version of the old audit log query
	query := "SELECT a.id, a.created_at, i.name, a.quantity, u.name, a.memo, a.type FROM activity_logs a LEFT JOIN items i ON a.item_id = i.id LEFT JOIN users u ON a.user_id = u.id ORDER BY a.created_at DESC"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var id int64
		var createdAt, itemName, userName, memo, logType string
		var quantity int
		if err := rows.Scan(&id, &createdAt, &itemName, &quantity, &userName, &memo, &logType); err != nil {
			continue
		}
		results = append(results, map[string]interface{}{
			"id":   id,
			"time": createdAt,
			"item": itemName,
			"qty":  quantity,
			"user": userName,
			"memo": memo,
			"type": logType,
		})
	}
	return results, nil
}
