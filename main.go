package main

import (
	"github.com/dchf12/jwt/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Setup(app)
	app.Listen("localhost:8000")
}
