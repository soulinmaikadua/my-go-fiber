package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soulinmaikadua/my-go-fiber/pkg/controllers"
)

func PostRoutes(app *fiber.App) {
	// create routes

	route := app.Group("/posts")

	route.Get("/", controllers.GetPosts)
}
