package controllers

import (
	"path/filepath"

	"github.com/dinever/golf"
)

func Initialize(app *golf.Application) *golf.Application {
	app.View.SetTemplateLoader("base", "app/views")
	app.View.SetTemplateLoader("admin", filepath.Join("app", "views", "admin"))
	app.View.SetTemplateLoader("theme", filepath.Join("app", "views", "default"))
	app.Static("/static/", "app/assets")
	app.SessionManager = golf.NewMemorySessionManager()

	app.Error(404, NotFoundController)
	registerAdminController(app)
	registerHomeController(app)

	return app
}

func registerHomeController(app *golf.Application) {
	statsChain := golf.NewChain()
	app.Get("/", statsChain.Final(HomeController))
}

func registerAdminController(app *golf.Application) {
	app.Get("/sign_up/", AuthSignUpGetController)
	app.Post("/sign_up/", AuthSignUpPostController)
	app.Get("/login/", AuthLoginGetController)
	app.Post("/login/", AuthLoginPostController)
}
