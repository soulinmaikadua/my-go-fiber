package controllers

import "github.com/gofiber/fiber/v2"


func GetUsers(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{
		"error": false,
		"message": nil,
		"data": nil,
	})
}