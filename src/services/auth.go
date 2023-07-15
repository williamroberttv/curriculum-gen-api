package services

import (
	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {
	jwtToken := "123456789"

	return c.Status(200).JSON(jwtToken)
}