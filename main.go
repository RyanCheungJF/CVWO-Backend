package main

import (
	"fmt"
	"os"

	"github.com/RyanCheungJF/CVWO-Backend/database"
	"github.com/RyanCheungJF/CVWO-Backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialize connection to database
	database.Connect()

	app := fiber.New()

	// Allow frontend to receive cookie (for HTTPOnly Cookies)
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	// Initialize paths for the app for each API Call
	routes.Setup(app)

	port := os.Getenv("PORT")
	app.Listen(fmt.Sprintf(":%s", port))
}
