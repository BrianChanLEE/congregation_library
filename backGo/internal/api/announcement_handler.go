package api

import (
	"boock/backGo/internal/common/response"
	"boock/backGo/internal/logger"
	"boock/backGo/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
// ...

// AnnouncementHandler는 공지사항 관련 핸들러를 관리합니다.
type AnnouncementHandler struct {
	AnnService service.AnnouncementServiceInterface
}

// NewAnnouncementHandler는 새로운 AnnouncementHandler를 생성합니다.
func NewAnnouncementHandler(service service.AnnouncementServiceInterface) *AnnouncementHandler {
	return &AnnouncementHandler{AnnService: service}
}

// GetAnnouncements 공지사항 목록 조회
func (h *AnnouncementHandler) GetAnnouncements(c *gin.Context) {
	anns, err := h.AnnService.GetAllAnnouncements()
	if err != nil {
		logger.Log.Error("공지사항 조회 실패", "error", err)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "조회 실패", err.Error())
		return
	}
	c.JSON(http.StatusOK, anns)
}

// AddAnnouncement 공지사항 등록
func (h *AnnouncementHandler) AddAnnouncement(c *gin.Context) {
	var req struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		AuthorID int64  `json:"author_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "입력값 오류", err.Error())
		return
	}

	if err := h.AnnService.CreateAnnouncement(req.Title, req.Content, req.AuthorID); err != nil {
		logger.Log.Error("공지사항 등록 실패", "error", err)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "등록 실패", err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "공지사항 등록 성공"})
}
// ... existing methods ...

// UpdateAnnouncement 공지사항 수정
func (h *AnnouncementHandler) UpdateAnnouncement(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "잘못된 ID", err.Error())
		return
	}

	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "입력값 오류", err.Error())
		return
	}

	if err := h.AnnService.UpdateAnnouncement(id, req.Title, req.Content); err != nil {
		logger.Log.Error("공지사항 수정 실패", "error", err, "id", id)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "수정 실패", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "공지사항 수정 성공"})
}

// DeleteAnnouncement 공지사항 삭제
func (h *AnnouncementHandler) DeleteAnnouncement(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "INVALID_INPUT", "잘못된 ID", err.Error())
		return
	}

	if err := h.AnnService.DeleteAnnouncement(id); err != nil {
		logger.Log.Error("공지사항 삭제 실패", "error", err, "id", id)
		response.SendError(c, http.StatusInternalServerError, "DB_ERROR", "삭제 실패", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "공지사항 삭제 성공"})
}

