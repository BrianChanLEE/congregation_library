package api

import (
	"boock/backGo/internal/common/response"
	"boock/backGo/internal/logger"
	"boock/backGo/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SystemHandler struct {
	SystemService service.SystemServiceInterface
}

func NewSystemHandler(service service.SystemServiceInterface) *SystemHandler {
	return &SystemHandler{SystemService: service}
}

// GetSystemStatus 시스템 상태 모니터링
func (h *SystemHandler) GetSystemStatus(c *gin.Context) {
	logs, err := h.SystemService.GetSystemErrors()
	if err != nil {
		logger.Log.Error("시스템 상태 조회 실패", "error", err)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "조회 실패", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      "ok",
		"error_logs":  logs,
		"system_time": "2026-05-29 15:00:00", // 예시
	})
}
