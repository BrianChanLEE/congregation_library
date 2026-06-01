package repository

import (
	"database/sql"
)

type AdminRepositoryInterface interface {
	GetStats() (int, int, int, error)
}

type AdminRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (r *AdminRepository) GetStats() (int, int, int, error) {
	var totalItems, recentActivity, pendingUsers int

	// 총 품목 수
	err := r.db.QueryRow("SELECT COUNT(*) FROM items").Scan(&totalItems)
	if err != nil {
		return 0, 0, 0, err
	}

	// 최근 7일간 활동 수
	err = r.db.QueryRow("SELECT COUNT(*) FROM activity_logs WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)").Scan(&recentActivity)
	if err != nil {
		return 0, 0, 0, err
	}

	// 가입 대기 사용자 수
	err = r.db.QueryRow("SELECT COUNT(*) FROM users WHERE status = 'PENDING'").Scan(&pendingUsers)
	if err != nil {
		return 0, 0, 0, err
	}

	return totalItems, recentActivity, pendingUsers, nil
}
