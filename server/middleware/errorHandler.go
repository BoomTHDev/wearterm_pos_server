package middleware

import (
	"errors"
	"log"

	"github.com/BoomTHDev/wear-pos-server/pkg/custom"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		if err == nil {
			return ctx.Next()
		}

		appErr := &custom.AppError{}
		fiberErr := &fiber.Error{}

		if errors.As(err, &appErr) {
			if appErr.StatusCode >= 500 && appErr.Err != nil {
				log.Printf("Internal AppError: %v, Original Err: %v", appErr.Message, appErr.Err)
			} else if appErr.Err != nil {
				log.Printf("AppError: %v, Original Err: %v", appErr.Message, appErr.Err)
			} else {
				log.Printf("AppError: %v", appErr.Message)
			}

			return ctx.Status(appErr.StatusCode).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"code":    appErr.Code,
					"message": appErr.Message,
				},
			})
		}

		if errors.As(err, &fiberErr) {
			log.Printf("Fiber Error: Code=%d, Message=%s", fiberErr.Code, fiberErr.Message)
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"code":    "FIBER_ERROR",
					"message": fiberErr.Message,
				},
			})
		}

		log.Printf("Unhandled Error: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error": fiber.Map{
				"code":    "INTERNAL_SERVER_ERROR",
				"message": "An unexpected internal server error occurred.",
			},
		})
	}
}
