package api_test

import (
	"boock/backGo/internal/logger"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 로거 초기화
	logger.Init()
	
	// 테스트 실행
	code := m.Run()
	
	os.Exit(code)
}
