package homepage

import (
	"reseller-jh-be/internal"
	"reseller-jh-be/internal/homepage/handler"
	"reseller-jh-be/internal/homepage/repository"
	"reseller-jh-be/internal/homepage/service"
)

func RegisterRoute(app *internal.Application) {
	homepageRepo := repository.NewHomepageRepository(app.DB.Postgres)
	homepageService := service.NewHomepageService(homepageRepo)
	homepageHandler := handler.NewHomepageHandler(homepageService)

	// Define user-related routes
	homepageRoutes := app.Gin.Group("/api")
	{
		homepageRoutes.GET("/homepage", homepageHandler.GetHomepage)
		homepageRoutes.PUT("/homepage", homepageHandler.UpdateHomepage)
	}
}
