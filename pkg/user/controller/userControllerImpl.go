package controller

import (
	"strconv"

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

// List godoc
// @Summary List all users
// @Description List all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {array} []model.UserResponse
// @Router /v1/users [get]
func (c *userControllerImpl) List(ctx *fiber.Ctx) error {
	users, appErr := c.userService.List()
	if appErr != nil {
		return appErr
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    users,
		"message": "User list successfully",
	})
}

// Read godoc
// @Summary Read user by id
// @Description Read user by id
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.UserResponse
// @Router /v1/users/{id} [get]
func (c *userControllerImpl) Read(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return custom.ErrInvalidInput("INVALID_REQUEST_BODY", "Failed to parse request body.", err)
	}

	user, appErr := c.userService.Read(uint64(id))
	if appErr != nil {
		return appErr
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
		"message": "User read successfully",
	})
}
