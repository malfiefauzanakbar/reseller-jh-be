package user

import (
	"reseller-jh-be/internal"
	"reseller-jh-be/internal/user/handler"
	"reseller-jh-be/internal/user/repository"
	"reseller-jh-be/internal/user/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func RegisterRoute(app *internal.Application) {
	userRepo := repository.NewUserRepository(app.DB.Postgres)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Define user-related routes
	userRoutes := app.Gin.Group("/api")
	{
		store := cookie.NewStore([]byte("secret"))
		userRoutes.Use(sessions.Sessions("mysession", store))
		userRoutes.POST("/register", userHandler.CreateUser)
		userRoutes.POST("/login", userHandler.Login)
	}
}
