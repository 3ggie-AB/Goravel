package routes

import (

	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"

	"goravel/app/helpers"

	"github.com/goravel/framework/contracts/http"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/api", func(ctx http.Context) http.Response {
		return ctx.Response().Json(200, http.Json{
			"message": helpers.Asset("public/logo.png"),
		})
	})
	facades.Route().Get("/users", userController.Index)      // Ambil semua user
	facades.Route().Get("/user/{id}", userController.Show)   // Ambil user berdasarkan ID
	facades.Route().Post("/user", userController.Store)      // Tambah user baru
	facades.Route().Put("/user/{id}", userController.Update) // Edit user berdasarkan ID
	facades.Route().Delete("/user/{id}", userController.Destroy)
}
