package main

import (
	"boock/backGo/internal/db"
	"boock/backGo/internal/http"
	"boock/backGo/internal/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	_ = godotenv.Load()

	logger.Init()
	logger.Log.Info("서버 시작 중...")

	requiredEnv := []string{
		"DB_USER",
		"DB_PASSWORD",
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
		"JWT_SECRET",
	}
	for _, key := range requiredEnv {
		if os.Getenv(key) == "" {
			log.Fatalf("필수 환경 변수가 누락되었습니다: %s", key)
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	sqlDB, err := db.Open(dsn)
	if err != nil {
		log.Fatalf("DB 연결 실패: %v", err)
	}
	defer sqlDB.Close()
	if err := db.EnsureSchema(sqlDB); err != nil {
		log.Fatalf("DB 스키마 점검/보정 실패: %v", err)
	}

	ginWriter := &logger.GinWriter{}
	gin.DefaultWriter = ginWriter
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		fmt.Fprintf(ginWriter, "\033[33m[GIN-debug] %-6s %-25s --> %s (%d handlers)\033[0m\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Initialize and Setup Routes
	h := http.InitializeHandlers(sqlDB)
	http.SetupRoutes(r, h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Log.Info(fmt.Sprintf("서버가 %s 포트에서 실행 중입니다.", port))
	r.Run(":" + port)
}
