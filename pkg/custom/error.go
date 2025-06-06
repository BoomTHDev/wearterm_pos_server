package custom

import "github.com/gofiber/fiber/v2"

func Error(ctx *fiber.Ctx, err error) error {
	if fe, ok := err.(*fiber.Error); ok {
		return ctx.Status(fe.Code).JSON(fiber.Map{
			"message": fe.Message,
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": fiber.ErrInternalServerError.Message,
	})
}
