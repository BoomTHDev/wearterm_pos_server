package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	List(c *fiber.Ctx) error
	Read(c *fiber.Ctx) error
}
