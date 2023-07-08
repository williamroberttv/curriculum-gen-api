package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type Credentials struct {
	password  string `json:"password"`
	email       string    `json:"email"`
}

func Authenticate(c *fiber.Ctx) error {
	jwtToken := "123456789"

	return c.Status(200).JSON(jwtToken)
}