package db

import (
	"database/sql"
	"fmt"
)

func ensureTable(db *sql.DB, table string, createSQL string) error {
	const q = `
SELECT 1
FROM information_schema.TABLES
WHERE TABLE_SCHEMA = DATABASE()
  AND TABLE_NAME = ?
LIMIT 1`

	var one int
	err := db.QueryRow(q, table).Scan(&one)
	if err == nil {
		return nil
	}
	if err != sql.ErrNoRows {
		return fmt.Errorf("테이블 점검 실패(table=%s): %w", table, err)
	}

	if _, err := db.Exec(createSQL); err != nil {
		return fmt.Errorf("테이블 생성 실패(table=%s): %w", table, err)
	}
	return nil
}

func ensureColumn(db *sql.DB, table string, column string, alterSQL string) error {
	const q = `
SELECT 1
FROM information_schema.COLUMNS
WHERE TABLE_SCHEMA = DATABASE()
  AND TABLE_NAME = ?
  AND COLUMN_NAME = ?
LIMIT 1`

	var one int
	err := db.QueryRow(q, table, column).Scan(&one)
	if err == nil {
		return nil
	}
	if err != sql.ErrNoRows {
		return fmt.Errorf("스키마 점검 실패(table=%s col=%s): %w", table, column, err)
	}

	if _, err := db.Exec(alterSQL); err != nil {
		return fmt.Errorf("스키마 보정 실패(table=%s col=%s): %w", table, column, err)
	}
	return nil
}

// EnsureSchema는 런타임 DB 스키마가 코드 기대치와 다를 때(특히 legacy DB) 최소한의 보정 작업을 수행합니다.
// 주 목적: soft delete 컬럼(`deleted_at`) 누락으로 인한 런타임 500 방지.
func EnsureSchema(db *sql.DB) error {
	if err := ensureTable(db, "activity_logs", `
CREATE TABLE IF NOT EXISTS activity_logs (
  id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT UNSIGNED NULL,
  item_id BIGINT UNSIGNED NOT NULL,
  quantity INT NOT NULL,
  type ENUM('IN','OUT','CANCEL') NOT NULL,
  method ENUM('WEB','QR') NOT NULL,
  memo TEXT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
)`); err != nil {
		return err
	}

	if err := ensureColumn(db, "items", "deleted_at", "ALTER TABLE `items` ADD COLUMN `deleted_at` DATETIME NULL"); err != nil {
		return err
	}
	if err := ensureColumn(db, "users", "deleted_at", "ALTER TABLE `users` ADD COLUMN `deleted_at` DATETIME NULL"); err != nil {
		return err
	}
	if err := ensureColumn(db, "announcements", "deleted_at", "ALTER TABLE `announcements` ADD COLUMN `deleted_at` DATETIME NULL"); err != nil {
		return err
	}
	return nil
}
