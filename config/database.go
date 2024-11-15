package config

import "os"

type DBConfig struct {
	User     string
	Password string
	Driver   string
	Name     string
	Host     string
	Port     string
}

type DB struct {
	Postgres DBConfig
}

func LoadDbConfig() DB {
	return DB{
		Postgres: DBConfig{
			User:     os.Getenv("POSTGRES_DB_USER"),
			Password: os.Getenv("POSTGRES_DB_PASSWORD"),
			Driver:   os.Getenv("POSTGRES_DB_DRIVER"),
			Name:     os.Getenv("POSTGRES_DB_NAME"),
			Host:     os.Getenv("POSTGRES_DB_HOST"),
			Port:     os.Getenv("POSTGRES_DB_PORT"),
		},
	}
}
