package routes

import (
	"github.com/RyanCheungJF/CVWO-Backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Server Status
	app.Get("", controllers.Status)

	// User Auth
	app.Post("api/register", controllers.Register)
	app.Post("api/login", controllers.Login)
	app.Get("api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)

	// Task Management
	app.Get("api/task/:userid", controllers.GetTask)
	app.Post("api/task", controllers.AddTask)
	app.Put("api/task/:id", controllers.UpdateTask)
	app.Delete("api/task/:id", controllers.DeleteTask)
}
