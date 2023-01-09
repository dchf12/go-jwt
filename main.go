package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/dchf12/jwt/routes"
)

func main() {
	app := fiber.New()
	routes.Setup(app)
	app.Listen("localhost:8000")
}
