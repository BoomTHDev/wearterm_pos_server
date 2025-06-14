package controller

import "github.com/gofiber/fiber/v2"

type ShopController interface {
	NewShop(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
}
