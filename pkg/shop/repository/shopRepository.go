package repository

import "github.com/BoomTHDev/wear-pos-server/entities"

type ShopRepository interface {
	CreateShop(shop *entities.Shop) (*entities.Shop, error)
	CheckShopExists(userID uint64, shopName string) (bool, error)
	ListShops(userID uint64) ([]entities.Shop, error)
	// UpdateShop()
	// DeleteShop()
}
