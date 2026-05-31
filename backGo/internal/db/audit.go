package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
)

// LogAction은 감사 로그에 기록할 작업 유형을 정의합니다.
func LogAction(ctx context.Context, db *sql.DB, userID int, action string, targetTable string, targetID int, details interface{}) error {
	// Note: 상세 내용을 JSON으로 직렬화합니다.
	detailsJSON, err := json.Marshal(details)
	if err != nil {
		return fmt.Errorf("로그 상세 정보 직렬화 실패: %w", err)
	}

	// Note: 감사 로그 테이블에 기록을 추가합니다.
	query := `INSERT INTO audit_logs (user_id, action, target_table, target_id, details) VALUES (?, ?, ?, ?, ?)`
	_, err = db.ExecContext(ctx, query, userID, action, targetTable, targetID, string(detailsJSON))
	if err != nil {
		return fmt.Errorf("감사 로그 저장 실패: %w", err)
	}

	return nil
}
