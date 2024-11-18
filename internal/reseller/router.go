package reseller

import (
	"reseller-jh-be/internal"
	"reseller-jh-be/internal/reseller/handler"
	"reseller-jh-be/internal/reseller/repository"
	"reseller-jh-be/internal/reseller/service"
	"reseller-jh-be/pkg/common"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func RegisterRoute(app *internal.Application) {
	resellerRepo := repository.NewResellerRepository(app.DB.Postgres)
	resellerService := service.NewResellerService(resellerRepo)
	resellerHandler := handler.NewResellerHandler(resellerService)

	route := app.Gin.Group("/api")
	{
		route.POST("/reseller", resellerHandler.CreateReseller)
		route.GET("/count-reseller", resellerHandler.CountResellers)
		route.GET("/reseller-chart", resellerHandler.ResellersChart)

		resellerRoutes := route.Group("/reseller")
		store := cookie.NewStore([]byte("secret"))
		resellerRoutes.Use(sessions.Sessions("mysession", store))
		resellerRoutes.Use(common.AuthMiddleware())
		{
			resellerRoutes.GET("", resellerHandler.GetAllReseller)
			resellerRoutes.GET("/:id", resellerHandler.GetReseller)
			resellerRoutes.PUT("/:id", resellerHandler.UpdateReseller)
			resellerRoutes.PUT("/read/:id", resellerHandler.ReadReseller)
			resellerRoutes.DELETE("/:id", resellerHandler.DeleteReseller)
		}
	}
}
