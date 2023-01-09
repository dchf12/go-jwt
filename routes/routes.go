package routes

import (
	"github.com/dchf12/jwt/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/user")
	api.Get("/get-user", controllers.User)
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)
	api.Post("/logout", controllers.Logout)
}
