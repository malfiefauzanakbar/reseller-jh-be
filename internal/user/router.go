package user

import (
	"reseller-jh-be/internal"
	"reseller-jh-be/internal/user/handler"
	"reseller-jh-be/internal/user/repository"
	"reseller-jh-be/internal/user/service"
)

func RegisterRoute(app *internal.Application) {
	userRepo := repository.NewUserRepository(app.DB.Postgres)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Define user-related routes
	userRoutes := app.Gin.Group("/api")
	{
		userRoutes.POST("/register", userHandler.CreateUser)
		userRoutes.POST("/login", userHandler.Login)
	}
}
