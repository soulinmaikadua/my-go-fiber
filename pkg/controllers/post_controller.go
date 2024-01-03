package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soulinmaikadua/my-go-fiber/pkg/configs"
)

func GetPosts(c *fiber.Ctx) error {
	// create connection
	db, err := configs.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}
	// Get all posts.
	posts, err := db.GetPosts()
	if err != nil {
		// Return, if posts not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "posts were not found",
			"total":   0,
			"data":    nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"total":   len(posts),
		"data":    posts,
	})
}
