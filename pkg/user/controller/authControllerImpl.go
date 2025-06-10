package controller

import (
	"strconv"

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

func (c *authControllerImpl) NewPIN(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return custom.ErrInvalidInput("INVALID_REQUEST_BODY", "Failed to parse request body. Please ensure your user id is valid.", err)
	}

	newPin := _userModel.NewPINRequest{}
	if err := ctx.BodyParser(&newPin); err != nil {
		return custom.ErrInvalidInput("INVALID_REQUEST_BODY", "Failed to parse request body. Please ensure it's valid JSON.", err)
	}

	if appErr := c.authService.NewPIN(uint64(id), newPin.Pin); appErr != nil {
		return appErr
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User pin created successfully",
	})
}

func (c *authControllerImpl) LoginWithPassword(ctx *fiber.Ctx) error {
	req := _userModel.LoginWithPasswordRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return custom.ErrInvalidInput("INVALID_REQUEST_BODY", "Failed to parse request body. Please ensure it's valid JSON.", err)
	}
	loginResponse, appErr := c.authService.LoginWithPassword(req)
	if appErr != nil {
		return appErr
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    loginResponse,
		"message": "User logged in successfully",
	})
}

func (c *authControllerImpl) LoginWithPin(ctx *fiber.Ctx) error {
	req := _userModel.LoginWithPinRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return custom.ErrInvalidInput("INVALID_REQUEST_BODY", "Failed to parse request body. Please ensure it's valid JSON.", err)
	}
	loginResponse, appErr := c.authService.LoginWithPin(req)
	if appErr != nil {
		return appErr
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    loginResponse,
		"message": "User logged in successfully",
	})
}
