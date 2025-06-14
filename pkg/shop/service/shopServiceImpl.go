package service

import (
	"github.com/BoomTHDev/wear-pos-server/entities"
	"github.com/BoomTHDev/wear-pos-server/pkg/custom"
	_shopModel "github.com/BoomTHDev/wear-pos-server/pkg/shop/model"
	_shopRepository "github.com/BoomTHDev/wear-pos-server/pkg/shop/repository"
)

type shopServiceImpl struct {
	shopRepository _shopRepository.ShopRepository
}

func NewShopServiceImpl(shopRepository _shopRepository.ShopRepository) ShopService {
	return &shopServiceImpl{shopRepository: shopRepository}
}

func (s *shopServiceImpl) NewShop(req _shopModel.NewShopRequest) (*_shopModel.ShopResponse, *custom.AppError) {
	if req.Name == "" {
		return nil, custom.ErrInvalidInput("INVALID_REQUEST_BODY", "Name is empty", nil)
	}

	if req.UserID <= 0 {
		return nil, custom.ErrInvalidInput("INVALID_REQUEST_BODY", "User ID is invalid", nil)
	}

	// Check if shop with the same name already exists for this user
	exists, err := s.shopRepository.CheckShopExists(req.UserID, req.Name)
	if err != nil {
		return nil, custom.ErrIntervalServer("INTERVAL_SERVER", "Failed to check shop availability", err)
	}

	if exists {
		return nil, custom.ErrConflict("SHOP_DUPLICATE", "You already have a shop with this name", nil)
	}

	shop := entities.Shop{
		Name:   req.Name,
		UserID: req.UserID,
	}

	newShop, err := s.shopRepository.CreateShop(&shop)
	if err != nil {
		return nil, custom.ErrIntervalServer("INTERVAL_SERVER", "Failed to create shop", err)
	}

	return _shopModel.ToShopResponse(newShop), nil
}
