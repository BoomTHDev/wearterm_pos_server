package controller

import (
	"github.com/BoomTHDev/wear-pos-server/pkg/custom"
	_userModel "github.com/BoomTHDev/wear-pos-server/pkg/user/model"
	_authService "github.com/BoomTHDev/wear-pos-server/pkg/user/service"
	"github.com/gofiber/fiber/v2"
)

type authControllerImpl struct {
	authService _authService.AuthService
}

func NewAuthController(authService _authService.AuthService) AuthController {
	return &authControllerImpl{authService: authService}
}

func (c *authControllerImpl) Register(ctx *fiber.Ctx) error {
	req := _userModel.RegisterRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return custom.ErrInvalidInput("INVALID_REQUEST_BODY", "Failed to parse request body. Please ensure it's valid JSON.", err)
	}
	user, appErr := c.authService.Register(req)
	if appErr != nil {
		return appErr
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    user,
		"message": "User registered successfully",
	})
}
