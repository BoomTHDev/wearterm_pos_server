package controller

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	Register(c *fiber.Ctx) error
	NewPIN(c *fiber.Ctx) error
	LoginWithPassword(c *fiber.Ctx) error
	LoginWithPin(c *fiber.Ctx) error
}
