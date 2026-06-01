package repository

import (
	"database/sql"
)

type SystemRepositoryInterface interface {
	GetErrorLogs() ([]map[string]interface{}, error)
}

type SystemRepository struct {
	db *sql.DB
}

func NewSystemRepository(db *sql.DB) *SystemRepository {
	return &SystemRepository{db: db}
}

func (r *SystemRepository) GetErrorLogs() ([]map[string]interface{}, error) {
	// 최근 발생한 에러 로그 (간단한 예시: 로그 테이블이 있다면 조회, 현재는 에러가 발생한 트랜잭션 등을 예시로 조회)
	// 실제 로그 파일 분석이 필요하다면 logger 패키지를 활용해야 함
	query := "SELECT id, memo, created_at FROM activity_logs WHERE memo LIKE '%error%' ORDER BY created_at DESC LIMIT 10"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []map[string]interface{}
	for rows.Next() {
		var id int64
		var memo, createdAt string
		if err := rows.Scan(&id, &memo, &createdAt); err != nil {
			return nil, err
		}
		logs = append(logs, map[string]interface{}{"id": id, "error": memo, "time": createdAt})
	}
	return logs, nil
}
