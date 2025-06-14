package service

import (
	"github.com/BoomTHDev/wear-pos-server/pkg/custom"
	_shopModel "github.com/BoomTHDev/wear-pos-server/pkg/shop/model"
)

type ShopService interface {
	NewShop(req _shopModel.NewShopRequest) (*_shopModel.ShopResponse, *custom.AppError)
	ListShops(userID uint64) ([]_shopModel.ShopResponses, *custom.AppError)
}
