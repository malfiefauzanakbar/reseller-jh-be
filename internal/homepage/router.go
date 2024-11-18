package homepage

import (
	"reseller-jh-be/internal"
	"reseller-jh-be/internal/homepage/handler"
	"reseller-jh-be/internal/homepage/repository"
	"reseller-jh-be/internal/homepage/service"
	"reseller-jh-be/pkg/common"
	
)

func RegisterRoute(app *internal.Application) {
	homepageRepo := repository.NewHomepageRepository(app.DB.Postgres)
	homepageService := service.NewHomepageService(homepageRepo)
	homepageHandler := handler.NewHomepageHandler(homepageService)
	
	route := app.Gin.Group("/api")
	{
		route.GET("/homepage", homepageHandler.GetHomepage)
		
		homepageRoutes := route.Group("/homepage")		
		homepageRoutes.Use(common.AuthMiddleware())
		{			
			homepageRoutes.PUT("", homepageHandler.UpdateHomepage)
		}		
	}
}
