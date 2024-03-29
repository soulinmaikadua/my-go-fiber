package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/soulinmaikadua/my-go-fiber/pkg/routes"
)

// var secretKey = []byte("your-secret-key")

func main() {
	app := fiber.New()

	// Load variables from .env file
	godotenv.Load()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello world!",
		})
	})

	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: secretKey},
	// }))

	// Routes
	routes.AuthRoutes(app)
	routes.UserRoutes(app)
	routes.PostRoutes(app)
	routes.NotFoundRoute(app)

	port := 6000
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is running on http://localhost:%d\n", port)

	err := app.Listen(addr)
	if err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}
