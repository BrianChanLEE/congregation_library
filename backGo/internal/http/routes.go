package http

import (
	"boock/backGo/internal/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes 는 모든 API 경로를 설정합니다.
func SetupRoutes(r *gin.Engine, h *HandlerContainer) {
	// 1. 공개 라우트 (인증 불필요)
	r.GET("/api/catalog", h.Catalog.GetCatalog)
	r.GET("/api/catalog/categories", h.Catalog.GetCategories)
	r.POST("/api/auth/register", h.Auth.Register)
	r.POST("/api/auth/login", h.Auth.Login)

	// 2. 인증 필요 라우트 (사용자/관리자 공통)
	authGroup := r.Group("/api")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.GET("/user/profile", h.User.GetProfile)
		authGroup.PUT("/user/password", h.User.ChangePassword)
		authGroup.POST("/transactions", h.Transaction.AddTransaction)
		authGroup.GET("/history", h.Transaction.GetAuditLogs)
		authGroup.GET("/inventory", h.Item.GetInventory)
	}

	// 3. 관리자 전용 라우트
	adminGroup := r.Group("/api/admin")
	adminGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminGroup.GET("/alerts", h.Admin.GetAlerts)
		adminGroup.POST("/items", h.Item.AddItem)
		adminGroup.GET("/items", h.Item.GetItems)
		adminGroup.PUT("/items/:id", h.Item.UpdateItem)
		adminGroup.DELETE("/items/:id", h.Item.DeleteItem)
		adminGroup.GET("/stats", h.Admin.GetAdminStats)
		adminGroup.GET("/users/pending", h.Admin.GetPendingUsers)
		adminGroup.PUT("/users/:id/status", h.Admin.UpdateUserStatus)
		adminGroup.DELETE("/users/:id", h.Admin.DeleteUser)
		adminGroup.GET("/announcements", h.Announcement.GetAnnouncements)
		adminGroup.POST("/announcements", h.Announcement.AddAnnouncement)
		adminGroup.PUT("/announcements/:id", h.Announcement.UpdateAnnouncement)
		adminGroup.DELETE("/announcements/:id", h.Announcement.DeleteAnnouncement)
		adminGroup.GET("/system-status", h.System.GetSystemStatus)
	}
}
