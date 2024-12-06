package internal

import (
	"reseller-jh-be/config"
	databases "reseller-jh-be/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Gin    *gin.Engine
	Config *config.Config
	DB     *databases.Connection
}

func NewApp(cfg *config.Config) *Application {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Static("/uploads", "./uploads")
	r.Use(sessions.Sessions("session", store))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://reseller.jimshoneyofficial.co.id"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	return &Application{
		Gin:    r,
		Config: cfg,
		DB:     databases.InitConnection(cfg),
	}
}

func (a *Application) Start() {
	a.Gin.Run(":" + a.Config.App.Port)
}
