package main

import (
	"reseller-jh-be/config"
	"reseller-jh-be/database/migration"
	"reseller-jh-be/internal"
	"reseller-jh-be/router"
)

func main() {
	app := internal.NewApp(config.NewConfig())

	//Migrate only when needed
	migration.MigrateDB(app.DB.Postgres)

	//Router
	router.ConfigureRoute(app)
	app.Start()
}
