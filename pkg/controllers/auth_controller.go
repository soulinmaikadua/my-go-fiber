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

func SignUp(c *fiber.Ctx) error {

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
	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	// Create a new validator for a User model.
	validate := validator.New()

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
	fmt.Println(user)

	// Validate user fields.
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": utils.ValidatorErrors(err),
		})
	}

	// Delete user by given ID.
	if err := db.SignUp(user); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	user.Password = ""
	token, err := utils.GenerateNewToken(*user)
	if err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	c.Set("token", token)
	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    user,
	})
}

func SingIn(c *fiber.Ctx) error {
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
	// Delete user by given ID.
	if err := db.SignUp(user); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	user.Password = ""
	token, err := utils.GenerateNewToken(*user)
	if err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	c.Set("token", token)
	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"message": nil,
		"data":    user,
	})
}
