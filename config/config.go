package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	App AppConfig
	DB  DB
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		DB:  LoadDbConfig(),
		App: LoadAppConfig(),
	}
}
