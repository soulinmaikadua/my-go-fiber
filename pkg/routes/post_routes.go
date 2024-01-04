package routes

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/soulinmaikadua/my-go-fiber/pkg/controllers"
)

func PostRoutes(app *fiber.App) {
	// create routes
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("your-secret-key")},
	}))
	route := app.Group("/posts")

	route.Get("/", controllers.GetPosts)
	route.Get("/:id", controllers.GetPost)
	route.Post("/", controllers.CreatePost)
	route.Put("/:id", controllers.UpdatePost)
}
