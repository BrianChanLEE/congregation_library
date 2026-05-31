package http

import (
	"boock/backGo/internal/api"
	"boock/backGo/internal/repository"
	"boock/backGo/internal/service"
)

type HandlerContainer struct {
	Transaction  *api.TransactionHandler
	Announcement *api.AnnouncementHandler
	Auth         *api.AuthHandler
	Catalog      *api.CatalogHandler
	Item         *api.ItemHandler
	Admin        *api.AdminHandler
	System       *api.SystemHandler
	User         *api.UserHandler
}

func InitializeHandlers() *HandlerContainer {
	activityRepo := &repository.ActivityLogRepository{}
	activityService := service.NewActivityLogService(activityRepo)
	
	announcementRepo := &repository.AnnouncementRepository{}
	announcementService := service.NewAnnouncementService(announcementRepo)
	
	userRepo := &repository.UserRepository{}
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)

	itemRepo := &repository.ItemRepository{}
	catalogService := service.NewCatalogService(itemRepo)
	itemService := service.NewItemService(itemRepo)

	adminRepo := &repository.AdminRepository{}
	adminService := service.NewAdminService(adminRepo)

	systemRepo := &repository.SystemRepository{}
	systemService := service.NewSystemService(systemRepo)

	return &HandlerContainer{
		Transaction:  api.NewTransactionHandler(activityService),
		Announcement: api.NewAnnouncementHandler(announcementService),
		Auth:         api.NewAuthHandler(authService),
		Catalog:      api.NewCatalogHandler(catalogService),
		Item:         api.NewItemHandler(itemService),
		Admin:        api.NewAdminHandler(adminService, userService),
		System:       api.NewSystemHandler(systemService),
		User:         api.NewUserHandler(userService),
	}
}
