package main

import (
	"log"
	"role-based/api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.Routes(app)

	log.Fatal(app.Listen(":4000"))

}
