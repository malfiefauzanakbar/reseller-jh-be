package reseller

import (
	"reseller-jh-be/internal"
	"reseller-jh-be/internal/reseller/handler"
	"reseller-jh-be/internal/reseller/repository"
	"reseller-jh-be/internal/reseller/service"
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
			resellerRoutes.GET("/", resellerHandler.GetAllReseller)
			resellerRoutes.GET("/:id", resellerHandler.GetReseller)
			resellerRoutes.PUT("/:id", resellerHandler.UpdateReseller)
			resellerRoutes.PUT("/read/:id", resellerHandler.ReadReseller)
			resellerRoutes.DELETE("/:id", resellerHandler.DeleteReseller)
			resellerRoutes.GET("/export/excel", resellerHandler.ExportExcelResellers)
			resellerRoutes.GET("/dashboard/count", resellerHandler.CountResellers)
			resellerRoutes.GET("/dashboard/chart", resellerHandler.ResellersChart)
		}
	}
}
