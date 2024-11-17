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
	
	resellerRoutes := app.Gin.Group("/api")
	{
		// resellerRoutes.POST("/reseller", resellerHandler.CreateReseller)
		// resellerRoutes.GET("/reseller", resellerHandler.GetAllReseller)
		// resellerRoutes.GET("/reseller/:id", resellerHandler.GetReseller)
		// resellerRoutes.PUT("/reseller/:id", resellerHandler.UpdateReseller)
		// resellerRoutes.PUT("/reseller/read/:id", resellerHandler.ReadReseller)
		// resellerRoutes.DELETE("/reseller/:id", resellerHandler.DeleteReseller)
		// resellerRoutes.GET("/count-reseller", resellerHandler.CountResellers)
		// resellerRoutes.GET("/reseller-chart", resellerHandler.ResellersChart)
		resel := resellerRoutes.Group("/reseller")
		resel.Use(common.AuthMiddleware())
		{
			resel.GET("/", resellerHandler.GetAllReseller)
		}		
	}
}
