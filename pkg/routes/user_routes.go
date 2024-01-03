package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soulinmaikadua/my-go-fiber/pkg/controllers"
)

func UserRoutes(app *fiber.App) {
	// create routes

	route := app.Group("/users")

	route.Get("/", controllers.GetUsers)
	route.Get("/:id", controllers.GetUser)
	route.Post("/", controllers.CreateUser)
	route.Put("/:id", controllers.UpdateUser)
	route.Delete("/:id", controllers.DeleteUser)
}
