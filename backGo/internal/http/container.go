package http

import (
	"boock/backGo/internal/api"
	"boock/backGo/internal/repository"
	"boock/backGo/internal/service"
	"database/sql"
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

func InitializeHandlers(db *sql.DB) *HandlerContainer {
	activityRepo := repository.NewActivityLogRepository(db)
	activityService := service.NewActivityLogService(activityRepo)

	announcementRepo := repository.NewAnnouncementRepository(db)
	announcementService := service.NewAnnouncementService(announcementRepo)

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)

	itemRepo := repository.NewItemRepository(db)
	catalogService := service.NewCatalogService(itemRepo)
	itemService := service.NewItemService(itemRepo)

	adminRepo := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepo)

	systemRepo := repository.NewSystemRepository(db)
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
