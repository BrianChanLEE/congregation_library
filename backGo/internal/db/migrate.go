package db

import (
	"database/sql"
	"fmt"
)

type migration struct {
	id string
	up func(*sql.DB) error
}

func ensureMigrationTable(db *sql.DB) error {
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS schema_migrations (
  id VARCHAR(128) PRIMARY KEY,
  applied_at DATETIME DEFAULT CURRENT_TIMESTAMP
)`)
	if err != nil {
		return fmt.Errorf("마이그레이션 테이블 생성 실패: %w", err)
	}
	return nil
}

func isMigrationApplied(db *sql.DB, id string) (bool, error) {
	var one int
	err := db.QueryRow("SELECT 1 FROM schema_migrations WHERE id = ? LIMIT 1", id).Scan(&one)
	if err == nil {
		return true, nil
	}
	if err == sql.ErrNoRows {
		return false, nil
	}
	return false, fmt.Errorf("마이그레이션 이력 조회 실패(id=%s): %w", id, err)
}

func markMigrationApplied(db *sql.DB, id string) error {
	if _, err := db.Exec("INSERT INTO schema_migrations (id) VALUES (?)", id); err != nil {
		return fmt.Errorf("마이그레이션 이력 저장 실패(id=%s): %w", id, err)
	}
	return nil
}

func applyMigrations(db *sql.DB, migrations []migration) error {
	if err := ensureMigrationTable(db); err != nil {
		return err
	}

	for _, migration := range migrations {
		applied, err := isMigrationApplied(db, migration.id)
		if err != nil {
			return err
		}
		if applied {
			continue
		}
		if err := migration.up(db); err != nil {
			return fmt.Errorf("마이그레이션 적용 실패(id=%s): %w", migration.id, err)
		}
		if err := markMigrationApplied(db, migration.id); err != nil {
			return err
		}
	}
	return nil
}

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
// 각 보정은 schema_migrations에 적용 이력을 남깁니다.
func EnsureSchema(db *sql.DB) error {
	return applyMigrations(db, []migration{
		{
			id: "20260601_ensure_activity_logs",
			up: func(db *sql.DB) error {
				return ensureTable(db, "activity_logs", `
CREATE TABLE IF NOT EXISTS activity_logs (
  id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT UNSIGNED NULL,
  item_id BIGINT UNSIGNED NOT NULL,
  quantity INT NOT NULL,
  type ENUM('IN','OUT','CANCEL') NOT NULL,
  method ENUM('WEB','QR') NOT NULL,
  memo TEXT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
)`)
			},
		},
		{
			id: "20260601_add_items_deleted_at",
			up: func(db *sql.DB) error {
				return ensureColumn(db, "items", "deleted_at", "ALTER TABLE `items` ADD COLUMN `deleted_at` DATETIME NULL")
			},
		},
		{
			id: "20260601_add_users_deleted_at",
			up: func(db *sql.DB) error {
				return ensureColumn(db, "users", "deleted_at", "ALTER TABLE `users` ADD COLUMN `deleted_at` DATETIME NULL")
			},
		},
		{
			id: "20260601_add_announcements_deleted_at",
			up: func(db *sql.DB) error {
				return ensureColumn(db, "announcements", "deleted_at", "ALTER TABLE `announcements` ADD COLUMN `deleted_at` DATETIME NULL")
			},
		},
	})
}
