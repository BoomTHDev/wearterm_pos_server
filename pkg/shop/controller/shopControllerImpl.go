package controller

import (
	"strconv"

	"github.com/BoomTHDev/wear-pos-server/pkg/custom"
	_shopModel "github.com/BoomTHDev/wear-pos-server/pkg/shop/model"
	_shopService "github.com/BoomTHDev/wear-pos-server/pkg/shop/service"
	"github.com/gofiber/fiber/v2"
)

type shopControllerImpl struct {
	shopService _shopService.ShopService
}

func NewShopControllerImpl(shopService _shopService.ShopService) ShopController {
	return &shopControllerImpl{shopService: shopService}
}

func (c *shopControllerImpl) NewShop(ctx *fiber.Ctx) error {
	userId, err := strconv.Atoi(ctx.Params("userId"))
	if err != nil {
		return custom.ErrInvalidInput("INVALID_REQUEST_BODY", "Failed to parse request body. Please ensure it's valid JSON.", err)
	}

	req := _shopModel.NewShopRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return custom.ErrInvalidInput("INVALID_REQUEST_BODY", "Failed to parse request body. Please ensure it's valid JSON.", err)
	}

	req.UserID = uint64(userId)

	shopResponse, appErr := c.shopService.NewShop(req)
	if appErr != nil {
		return appErr
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    shopResponse,
		"message": "Shop created successfully",
	})
}
