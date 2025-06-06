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
		return custom.Error(ctx, err)
	}

	user, err := c.userService.Add(&userEntity)
	if err != nil {
		return custom.Error(ctx, err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}
