package api

import (
	"boock/backGo/internal/common/response"
	"boock/backGo/internal/logger"
	"boock/backGo/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserService service.UserServiceInterface
}

func NewUserHandler(service service.UserServiceInterface) *UserHandler {
	return &UserHandler{UserService: service}
}

// GetProfile 사용자 프로필 조회
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		response.SendError(c, http.StatusUnauthorized, "AUTH_REQUIRED", "인증이 필요합니다.", "")
		return
	}

	user, err := h.UserService.GetUserProfile(strconv.FormatInt(userID.(int64), 10))
	if err != nil {
		logger.Log.Error("프로필 조회 실패", "error", err, "userID", userID)
		response.SendError(c, http.StatusNotFound, "NOT_FOUND", "사용자를 찾을 수 없습니다.", err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// ChangePassword 비밀번호 변경
func (h *UserHandler) ChangePassword(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		response.SendError(c, http.StatusUnauthorized, "AUTH_REQUIRED", "인증이 필요합니다.", "")
		return
	}

	var req struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "입력값 오류", err.Error())
		return
	}

	if err := h.UserService.ChangePassword(userID.(int64), req.CurrentPassword, req.NewPassword); err != nil {
		logger.Log.Error("비밀번호 변경 실패", "error", err, "userID", userID)
		response.SendError(c, http.StatusInternalServerError, "UPDATE_ERROR", "비밀번호 변경 실패", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "비밀번호 변경 성공"})
}
