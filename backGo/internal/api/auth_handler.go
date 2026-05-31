package api

import (
	"boock/backGo/internal/common/response"
	"boock/backGo/internal/logger"
	"boock/backGo/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService service.AuthServiceInterface
}

func NewAuthHandler(service service.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{AuthService: service}
}

// Register 회원가입 핸들러
func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Warn("회원가입 요청 데이터 바인딩 실패", "error", err)
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "입력값을 확인해주세요.", err.Error())
		return
	}

	if err := h.AuthService.Register(req.Name, req.Password); err != nil {
		logger.Log.Error("회원가입 실패", "error", err, "user", req.Name)
		response.SendError(c, http.StatusInternalServerError, "AUTH_ERROR", "회원가입에 실패했습니다.", err.Error())
		return
	}

	logger.Log.Info("사용자 가입 성공", "user", req.Name)
	c.JSON(http.StatusCreated, gin.H{"message": "가입 성공"})
}

// Login 로그인 핸들러
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		CongCode   string `json:"cong_code" binding:"required"`
		JwhubEmail string `json:"jwhub_email" binding:"required"`
		Password   string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Warn("로그인 요청 데이터 바인딩 실패", "error", err)
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "입력값을 확인해주세요.", err.Error())
		return
	}

	token, err := h.AuthService.Login(req.CongCode, req.JwhubEmail, req.Password)
	if err != nil {
		logger.Log.Warn("로그인 실패", "email", req.JwhubEmail, "error", err)
		response.SendError(c, http.StatusUnauthorized, "AUTH_FAILED", "인증에 실패했습니다.", err.Error())
		return
	}

	logger.Log.Info("사용자 로그인 성공", "email", req.JwhubEmail)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
