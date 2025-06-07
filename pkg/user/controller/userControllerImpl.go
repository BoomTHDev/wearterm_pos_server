package controller

import (
	"github.com/BoomTHDev/wear-pos-server/entities"
	"github.com/BoomTHDev/wear-pos-server/pkg/custom"
	_userService "github.com/BoomTHDev/wear-pos-server/pkg/user/service"
	"github.com/gofiber/fiber/v2"
)

type userControllerImpl struct {
	userService _userService.UserService
}

func NewUserController(userService _userService.UserService) UserController {
	return &userControllerImpl{userService: userService}
}

func (c *userControllerImpl) Add(ctx *fiber.Ctx) error {
	userEntity := entities.User{}
	if err := ctx.BodyParser(&userEntity); err != nil {
		return custom.ErrInvalidInput("INVALID_REQUEST_BODY", "Failed to parse request body. Please ensure it's valid JSON.", err)
	}

	user, appErr := c.userService.Add(&userEntity)
	if appErr != nil {
		return appErr
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    user,
		"message": "User added successfully",
	})
}

func (c *userControllerImpl) List(ctx *fiber.Ctx) error {
	users, appErr := c.userService.List()
	if appErr != nil {
		return appErr
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": users,
		"message": "User list successfully"
	})
}