package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
	port := 5000
	fmt.Printf("Server is running on https://localhost:%d\n", port)

    app.Listen(fmt.Sprintf(":%d", port))
}