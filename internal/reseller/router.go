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

	app.Gin.POST("/api/reseller", resellerHandler.CreateReseller)
	route := app.Gin.Group("/api")
	route.Use(common.AuthMiddleware())
	{						
		route.GET("/reseller", resellerHandler.GetAllReseller)
		route.GET("/reseller/:id", resellerHandler.GetReseller)
		route.PUT("/reseller/:id", resellerHandler.UpdateReseller)
		route.PUT("/reseller/read/:id", resellerHandler.ReadReseller)
		route.DELETE("/reseller/:id", resellerHandler.DeleteReseller)
		route.GET("/reseller/export/excel", resellerHandler.ExportExcelResellers)
		route.GET("/reseller/dashboard/count", resellerHandler.CountResellers)
		route.GET("/reseller/dashboard/chart", resellerHandler.ResellersChart)		
	}
}
