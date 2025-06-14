package repository

import "github.com/BoomTHDev/wear-pos-server/entities"

type ShopRepository interface {
	CreateShop(shop *entities.Shop) (*entities.Shop, error)
	CheckShopExists(userID uint64, shopName string) (bool, error)
	// UpdateShop()
	// DeleteShop()
}
