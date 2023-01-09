package main

import (
	"github.com/dchf12/jwt/database"
	"github.com/dchf12/jwt/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DBconn()
	app := fiber.New()
	routes.Setup(app)
	app.Listen("localhost:8000")
}
