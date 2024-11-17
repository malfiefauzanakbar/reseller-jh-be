package router

import (
	"reseller-jh-be/internal"
	"reseller-jh-be/internal/reseller"
	"reseller-jh-be/internal/homepage"
	"reseller-jh-be/internal/user"
)

func ConfigureRoute(app *internal.Application) {
	user.RegisterRoute(app)
	reseller.RegisterRoute(app)
	homepage.RegisterRoute(app)
}
