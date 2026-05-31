package api

import (
	"boock/backGo/internal/common/response"
	"boock/backGo/internal/logger"
	"boock/backGo/internal/models"
	"boock/backGo/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ItemHandler struct {
	ItemService service.ItemServiceInterface
}

func NewItemHandler(service service.ItemServiceInterface) *ItemHandler {
	return &ItemHandler{ItemService: service}
}

// AddItem 신규 품목 등록 핸들러
func (h *ItemHandler) AddItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "입력값을 확인해주세요.", err.Error())
		return
	}

	if err := h.ItemService.AddItem(item.Name, item.Code); err != nil {
		logger.Log.Error("품목 등록 실패", "error", err, "item", item)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "품목 등록에 실패했습니다.", err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "품목 추가 성공"})
}

// GetItems 품목 목록 조회 핸들러
func (h *ItemHandler) GetItems(c *gin.Context) {
	items, err := h.ItemService.GetAllItems()
	if err != nil {
		logger.Log.Error("품목 조회 실패", "error", err)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "품목 조회에 실패했습니다.", err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}

// UpdateItem 품목 수정 핸들러
func (h *ItemHandler) UpdateItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "잘못된 ID", err.Error())
		return
	}
	var req struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "입력값 오류", err.Error())
		return
	}

	if err := h.ItemService.UpdateItem(id, req.Name, req.Code); err != nil {
		logger.Log.Error("품목 수정 실패", "error", err, "id", id)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "품목 수정 실패", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "품목 수정 성공"})
}

// DeleteItem 품목 삭제 핸들러 (Soft Delete)
func (h *ItemHandler) DeleteItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "잘못된 ID", err.Error())
		return
	}

	if err := h.ItemService.DeleteItem(id); err != nil {
		logger.Log.Error("품목 삭제 실패", "error", err, "id", id)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "품목 삭제 실패", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "품목 삭제 성공"})
}
