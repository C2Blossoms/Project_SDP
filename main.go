package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/register", Register)
	app.Post("/login", Login)
	app.Get("/protected", ProtectedRoute)

	app.Listen(":3000")
}
