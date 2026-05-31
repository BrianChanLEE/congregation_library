package api

import (
	"boock/backGo/internal/common/response"
	"boock/backGo/internal/logger"
	"boock/backGo/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// TransactionHandler는 거래 관련 핸들러를 관리합니다.
type TransactionHandler struct {
	ActivityService service.ActivityLogServiceInterface
}

// NewTransactionHandler는 새로운 TransactionHandler를 생성합니다.
func NewTransactionHandler(service service.ActivityLogServiceInterface) *TransactionHandler {
	return &TransactionHandler{ActivityService: service}
}

// AddTransaction 입고/출고 거래 생성 핸들러
func (h *TransactionHandler) AddTransaction(c *gin.Context) {
	var req struct {
		UserID   int64  `json:"user_id"`
		ItemID   int64  `json:"item_id"`
		Quantity int    `json:"quantity"`
		Type     string `json:"type"` // 'IN' or 'OUT'
		Method   string `json:"method"`
		Memo     string `json:"memo"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "입력값을 확인해주세요.", err.Error())
		return
	}

	if err := h.ActivityService.CreateLog(req.UserID, req.ItemID, req.Quantity, req.Type, req.Method, req.Memo); err != nil {
		logger.Log.Error("활동 로그 기록 실패", "error", err, "req", req)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "거래 생성에 실패했습니다.", err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "거래 성공"})
}

// CancelTransaction 트랜잭션 취소 핸들러
func (h *TransactionHandler) CancelTransaction(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "잘못된 ID", err.Error())
		return
	}

	if err := h.ActivityService.CancelLog(id); err != nil {
		logger.Log.Error("거래 취소 실패", "error", err, "id", id)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "취소 실패", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "취소 성공"})
}

// GetAuditLogs 전체 활동 내역 조회 핸들러
func (h *TransactionHandler) GetAuditLogs(c *gin.Context) {
	logs, err := h.ActivityService.GetAllLogs()
	if err != nil {
		logger.Log.Error("활동 내역 조회 실패", "error", err)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "활동 내역 조회에 실패했습니다.", err.Error())
		return
	}
	c.JSON(http.StatusOK, logs)
}

// GetOldAuditLogs 상세 활동 내역 조회 핸들러 (Old API 하위 호환)
func (h *TransactionHandler) GetOldAuditLogs(c *gin.Context) {
	logs, err := h.ActivityService.GetDetailedLogs()
	if err != nil {
		logger.Log.Error("상세 활동 내역 조회 실패", "error", err)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "조회 실패", err.Error())
		return
	}
	c.JSON(http.StatusOK, logs)
}
