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

// Register godoc
// @Summary Register new user
// @Description Register new user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body model.RegisterRequest true "Register Request"
// @Success 200 {object} model.UserResponse
// @Router /v1/auth/register [post]
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

// NewPIN godoc
// @Summary Create new pin for user
// @Description Create new pin for user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param body body model.NewPINRequest true "New PIN Request"
// @Success 200 {object} model.UserResponse
// @Router /v1/auth/new-pin/{id} [post]
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

// LoginWithPassword godoc
// @Summary Login with password
// @Description Login with password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body model.LoginWithPasswordRequest true "Login with password request"
// @Success 200 {object} model.LoginResponse
// @Router /v1/auth/login [post]
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

// LoginWithPin godoc
// @Summary Login with pin
// @Description Login with pin
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body model.LoginWithPinRequest true "Login with pin request"
// @Success 200 {object} model.LoginResponse
// @Router /v1/auth/login-with-pin [post]
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
