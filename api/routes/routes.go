package routes

import (
	"role-based/api/handlers"
	"role-based/config/database"
	"role-based/repository"
	"role-based/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	//"github.com/gofiber/fiber/v2/middleware/cors"
)

func Routes(app fiber.Router) {
	//Call Database Connection
	database.DbConnection()

	//Initializer
	repo := repository.AccountRepositoryInit(database.DB)
	service := services.AccountServicesInit(repo)
	handler := handlers.AccountHandlersInit(service)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Your Next.js frontend
		AllowCredentials: true,
	}))

	app.Post("/create", handler.CreateAccount)
	app.Post("/login", handler.AccountLogin)

}
