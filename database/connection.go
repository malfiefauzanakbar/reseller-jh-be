package databases

import (
	"fmt"

	"reseller-jh-be/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Connection struct {
	Postgres *gorm.DB
}

func InitConnection(cfg *config.Config) *Connection {
	return &Connection{
		Postgres: GetPostgresConnection(cfg),
	}
}

func GetPostgresConnection(cfg *config.Config) *gorm.DB {
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.DB.Postgres.Host,
		cfg.DB.Postgres.User,
		cfg.DB.Postgres.Password,
		cfg.DB.Postgres.Name,
		cfg.DB.Postgres.Port)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable SQL logging
	})
	if err != nil {
		panic(err.Error())
	}

	return db
}
