package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("DB 열기 실패: %v", err)
	}

	// 성능 최적화: 커넥션 풀 조정
	DB.SetMaxOpenConns(50) 
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)
	DB.SetConnMaxIdleTime(2 * time.Minute)

	if err = DB.Ping(); err != nil {
		log.Fatalf("DB 연결 실패: %v", err)
	}

	fmt.Println("DB 최적화 연결 성공")
}
