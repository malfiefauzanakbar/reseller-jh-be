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

	// Define user-related routes
	resellerRoutes := app.Gin.Group("/api")
	{
		resellerRoutes.POST("/reseller", resellerHandler.CreateReseller)
		resellerRoutes.GET("/reseller", resellerHandler.GetAllReseller)
		resellerRoutes.GET("/reseller/:id", resellerHandler.GetReseller)
		resellerRoutes.PUT("/reseller/:id", resellerHandler.UpdateReseller)
		resellerRoutes.PUT("/reseller/read/:id", resellerHandler.ReadReseller)
		resellerRoutes.DELETE("/reseller/:id", resellerHandler.DeleteReseller)
	}
}
