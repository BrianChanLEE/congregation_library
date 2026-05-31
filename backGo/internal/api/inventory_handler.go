package api

import (
	"boock/backGo/internal/common/response"
	"boock/backGo/internal/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetInventory 재고 현황 조회 핸들러
func (h *ItemHandler) GetInventory(c *gin.Context) {
	congIDStr := c.Query("cong_id")
	itemIDStr := c.Query("item_id")
	
	congID, _ := strconv.ParseInt(congIDStr, 10, 64)
	itemID, _ := strconv.ParseInt(itemIDStr, 10, 64)

	inventory, err := h.ItemService.GetInventory(congID, itemID)
	if err != nil {
		logger.Log.Error("재고 조회 실패", "error", err, "cong_id", congID, "item_id", itemID)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "재고 조회에 실패했습니다.", err.Error())
		return
	}
	c.JSON(http.StatusOK, inventory)
}
