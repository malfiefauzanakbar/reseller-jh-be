package router

import (
	"reseller-jh-be/internal"
	"reseller-jh-be/internal/reseller"
)

func ConfigureRoute(app *internal.Application) {
	reseller.RegisterRoute(app)
}
