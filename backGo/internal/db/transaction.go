package db

import (
	"database/sql"
	"fmt"
)

// TransactionFunc는 트랜잭션 내에서 실행될 함수 타입입니다.
type TransactionFunc func(*sql.Tx) error

// WithTransaction은 트랜잭션 시작, 커밋, 롤백을 자동으로 관리합니다.
func WithTransaction(db *sql.DB, fn TransactionFunc) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("트랜잭션 시작 실패: %w", err)
	}

	// Note: 패닉 발생 시 롤백
	defer tx.Rollback()

	if err := fn(tx); err != nil {
		return err
	}

	return tx.Commit()
}
