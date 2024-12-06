package reseller

import (
	"reseller-jh-be/internal"
	"reseller-jh-be/internal/reseller/handler"
	"reseller-jh-be/internal/reseller/repository"
	"reseller-jh-be/internal/reseller/service"
	"reseller-jh-be/pkg/common"
)

func RegisterRoute(app *internal.Application) {
	resellerRepo := repository.NewResellerRepository(app.DB.Postgres)
	resellerService := service.NewResellerService(resellerRepo)
	resellerHandler := handler.NewResellerHandler(resellerService)

	route := app.Gin.Group("/api")
	{
		resellerRoutes := route.Group("/reseller")
		// resellerRoutes.Use(common.AuthMiddleware())
		{
			resellerRoutes.POST("/", resellerHandler.CreateReseller)
			resellerRoutes.GET("/", common.AuthMiddleware(), resellerHandler.GetAllReseller)
			resellerRoutes.GET("/:id", common.AuthMiddleware(), resellerHandler.GetReseller)
			resellerRoutes.PUT("/:id", common.AuthMiddleware(), resellerHandler.UpdateReseller)
			resellerRoutes.PUT("/read/:id", common.AuthMiddleware(), resellerHandler.ReadReseller)
			resellerRoutes.DELETE("/:id", common.AuthMiddleware(), resellerHandler.DeleteReseller)
			resellerRoutes.GET("/export/excel", common.AuthMiddleware(), resellerHandler.ExportExcelResellers)
			resellerRoutes.GET("/dashboard/count", common.AuthMiddleware(), resellerHandler.CountResellers)
			resellerRoutes.GET("/dashboard/chart", common.AuthMiddleware(), resellerHandler.ResellersChart)
		}
	}
}
