package api

import (
	"boock/backGo/internal/common/response"
	"boock/backGo/internal/logger"
	"boock/backGo/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CatalogHandler struct {
	CatalogService service.CatalogServiceInterface
}

func NewCatalogHandler(service service.CatalogServiceInterface) *CatalogHandler {
	return &CatalogHandler{CatalogService: service}
}

// GetCatalog 전체 출판물 목록 조회 핸들러
func (h *CatalogHandler) GetCatalog(c *gin.Context) {
	items, err := h.CatalogService.GetCatalog()
	if err != nil {
		logger.Log.Error("카탈로그 조회 실패", "error", err)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "데이터 조회 중 오류가 발생했습니다.", err.Error())
		return
	}

	logger.Log.Info("카탈로그 목록 조회 성공", "count", len(items))
	c.JSON(http.StatusOK, items)
}

// GetCategories 카테고리 목록 조회 핸들러
func (h *CatalogHandler) GetCategories(c *gin.Context) {
	cats, err := h.CatalogService.GetCategories()
	if err != nil {
		logger.Log.Error("카테고리 조회 실패", "error", err)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "조회 실패", err.Error())
		return
	}
	c.JSON(http.StatusOK, cats)
}
