package main

import (
	"boock/backGo/internal/http"
	"boock/backGo/internal/db"
	"boock/backGo/internal/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	logger.Init()
	logger.Log.Info("서버 시작 중...")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db.InitDB(dsn)

	ginWriter := &logger.GinWriter{}
	gin.DefaultWriter = ginWriter
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		fmt.Fprintf(ginWriter, "\033[33m[GIN-debug] %-6s %-25s --> %s (%d handlers)\033[0m\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Initialize and Setup Routes
	h := http.InitializeHandlers()
	http.SetupRoutes(r, h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Log.Info(fmt.Sprintf("서버가 %s 포트에서 실행 중입니다.", port))
	r.Run(":" + port)
}
