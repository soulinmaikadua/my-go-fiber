package controllers

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/soulinmaikadua/my-go-fiber/pkg/configs"
	"github.com/soulinmaikadua/my-go-fiber/pkg/models"
	"github.com/soulinmaikadua/my-go-fiber/pkg/utils"
)

func GetPosts(c *fiber.Ctx) error {
	claims := utils.GetToken(c)
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
	parsedUUID, _ := uuid.Parse(claims["id"].(string))
	// Get all posts.
	posts, err := db.GetPosts(parsedUUID)
	if err != nil {
		// Return, if posts not found.
		fmt.Println(err.Error())
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

func GetPost(c *fiber.Ctx) error {
	// Catch user ID from URL.
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// Create database connection.
	db, err := configs.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// Get post by ID.
	post, err := db.GetPost(id)
	if err != nil {
		// Return, if post not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "post with the given ID is not found",
			"data":    nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    post,
	})
}

func CreatePost(c *fiber.Ctx) error {

	post := &models.Post{}
	claims := utils.GetToken(c)
	// Check, if received JSON data is valid.
	if err := c.BodyParser(post); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// Create database connection.
	db, err := configs.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// Create a new validator for a Post model.
	validate := validator.New()

	// Assign data
	post.ID = uuid.New()
	post.CreatedAt = time.Now()
	parsedUUID, _ := uuid.Parse(claims["id"].(string))
	post.UserId = parsedUUID

	// Validate post fields.
	if err := validate.Struct(post); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": utils.ValidatorErrors(err),
		})
	}
	// Delete post by given ID.
	if err := db.CreatePost(post); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    post,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	// Catch post ID from URL.
	id, err := uuid.Parse(c.Params("id"))
	claims := utils.GetToken(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// create new post struct
	post := &models.PostUpdate{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(post); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// Create database connection.
	db, err := configs.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// Checking, if user with given ID does exist.
	foundPost, err := db.GetPost(id)
	if err != nil {
		// Return status 404 and user not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with this ID not found",
		})
	}
	// Set initialized default data for user:
	post.UpdatedAt = time.Now()

	// Create a new validator for a user model.
	validate := utils.NewValidator()

	// Validate user fields.
	if err := validate.Struct(post); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": utils.ValidatorErrors(err),
		})
	}
	parsedUUID, _ := uuid.Parse(claims["id"].(string))
	postAndUserId := &models.PostAndUserId{
		ID:     foundPost.ID,
		UserId: parsedUUID,
	}
	// Update user by given ID.
	if err := db.UpdatePost(postAndUserId, post); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// Return status 200.
	return c.JSON(fiber.Map{
		"error":   false,
		"message": "success",
		"data":    nil,
	})

}

func DeletePost(c *fiber.Ctx) error {
	// Catch user ID from URL.
	id, err := uuid.Parse(c.Params("id"))
	claims := utils.GetToken(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// Create database connection.
	db, err := configs.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// Checking, if post with given ID does exist.
	foundPost, err := db.GetPost(id)
	if err != nil {
		// Return status 404 and post not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "post with this ID not found",
		})
	}

	parsedUUID, _ := uuid.Parse(claims["id"].(string))
	// Delete post by given ID.
	if err := db.DeletePost(foundPost.ID, parsedUUID); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	// Return status 204 no content.
	return c.JSON(fiber.Map{
		"error":   false,
		"message": "success",
	})
}
