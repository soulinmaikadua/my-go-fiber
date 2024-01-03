package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soulinmaikadua/my-go-fiber/pkg/controllers"
)

func AuthRoutes(app *fiber.App) {
	// create routes

	route := app.Group("/auth")

	route.Post("/signup", controllers.SignUp)
	route.Post("/signin", controllers.SingIn)
}
