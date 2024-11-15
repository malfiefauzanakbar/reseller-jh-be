package internal

import (
	"reseller-jh-be/config"
	databases "reseller-jh-be/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Gin    *gin.Engine
	Config *config.Config
	DB     *databases.Connection
}

func NewApp(cfg *config.Config) *Application {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
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
