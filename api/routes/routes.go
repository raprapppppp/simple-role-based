package routes

import (
	"role-based/api/handlers"
	"role-based/api/middleware"
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

	//Initializer for Accounts
	repo := repository.AccountRepositoryInit(database.DB)
	service := services.AccountServicesInit(repo)
	handler := handlers.AccountHandlersInit(service)


	//Init Task API
	taskRepo := repository.TaskRepoInit(database.DB)
	taskService := services.TaskServicesInit(taskRepo)
	taskHandler := handlers.TaskHandlerInit(taskService)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Your Next.js frontend
		AllowCredentials: true,
	}))

	app.Post("/create", handler.CreateAccount)
	app.Post("/login", handler.AccountLogin)

	taskGroup := app.Group("/task",middleware.RoleBasedMiddleware("admin","user"))
	taskGroup.Post("/create",taskHandler.CreateTask)
	taskGroup.Get("/get", taskHandler.GetTask)
	taskGroup.Get("/profile",  handler.GetProfile)

 
}
