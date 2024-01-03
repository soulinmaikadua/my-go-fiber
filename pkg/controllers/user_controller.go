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

func GetUsers(c *fiber.Ctx) error {
	fmt.Println("Hello world")
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
	// Get all users.
	users, err := db.GetUsers()
	if err != nil {
		// Return, if users not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "users were not found",
			"total":   0,
			"data":    nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"total":   len(users),
		"data":    users,
	})
}

func GetUser(c *fiber.Ctx) error {
	// Catch book ID from URL.
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
	// Get user by ID.
	user, err := db.GetUser(id)
	fmt.Println(user.Email)
	if err != nil {
		// Return, if user not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "user with the given ID is not found",
			"data":    nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    user,
	})
}

func CreateUser(c *fiber.Ctx) error {

	user := &models.User{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(user); err != nil {
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

	// Create a new validator for a User model.
	// validate := utils.NewValidator()
	validate := validator.New()

	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": utils.ValidatorErrors(err),
		})
	}
	user.Password = hashedPassword

	// Validate book fields.
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": utils.ValidatorErrors(err),
		})
	}
	// Delete book by given ID.
	if err := db.CreateUser(user); err != nil {
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
		"data":    user,
	})
}
