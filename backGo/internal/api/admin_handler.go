package api

import (
	"boock/backGo/internal/common/response"
	"boock/backGo/internal/logger"
	"boock/backGo/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	AdminService service.AdminServiceInterface
	UserService  service.UserServiceInterface
}

func NewAdminHandler(adminService service.AdminServiceInterface, userService service.UserServiceInterface) *AdminHandler {
	return &AdminHandler{AdminService: adminService, UserService: userService}
}

// GetAdminStats 대시보드 통계 조회
func (h *AdminHandler) GetAdminStats(c *gin.Context) {
	totalItems, recentActivity, pendingUsers, err := h.AdminService.GetStats()
	if err != nil {
		logger.Log.Error("통계 조회 실패", "error", err)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "통계 조회 실패", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total_items":           totalItems,
		"recent_activity_count": recentActivity,
		"pending_user_count":    pendingUsers,
	})
}

// GetAlerts 시스템 알림 조회 (예시)
func (h *AdminHandler) GetAlerts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"alerts": []string{}})
}

// GetPendingUsers 가입 대기 사용자 목록 조회
func (h *AdminHandler) GetPendingUsers(c *gin.Context) {
	users, err := h.UserService.GetPendingUsers()
	if err != nil {
		logger.Log.Error("대기 사용자 조회 실패", "error", err)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "조회 실패", err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

// UpdateUserStatus 사용자 승인 상태 변경
func (h *AdminHandler) UpdateUserStatus(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "잘못된 ID", err.Error())
		return
	}

	var req struct {
		Status string `json:"status"` // 'APPROVED' or 'REJECTED'
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "입력값 오류", err.Error())
		return
	}

	if err := h.UserService.UpdateUserStatus(userID, req.Status); err != nil {
		logger.Log.Error("사용자 상태 변경 실패", "error", err, "userID", userID)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "상태 변경 실패", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "상태 변경 성공"})
}

// DeleteUser 사용자 삭제 (Soft Delete)
func (h *AdminHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "잘못된 ID", err.Error())
		return
	}

	if err := h.UserService.DeleteUser(userID); err != nil {
		logger.Log.Error("사용자 삭제 실패", "error", err, "userID", userID)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "사용자 삭제 실패", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "사용자 삭제 성공"})
}
