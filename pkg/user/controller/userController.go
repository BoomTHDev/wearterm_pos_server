package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	Add(c *fiber.Ctx) error
}
