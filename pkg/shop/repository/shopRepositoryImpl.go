package repository

import (
	"github.com/BoomTHDev/wear-pos-server/databases"
	"github.com/BoomTHDev/wear-pos-server/entities"
)

type shopRepositoryImpl struct {
	db databases.Database
}

func NewShopRepositoryImpl(db databases.Database) ShopRepository {
	return &shopRepositoryImpl{db: db}
}

func (r *shopRepositoryImpl) CreateShop(shop *entities.Shop) (*entities.Shop, error) {
	conn := r.db.ConnectionGetting()

	if err := conn.Create(&shop).Error; err != nil {
		return nil, err
	}

	return shop, nil
}

func (r *shopRepositoryImpl) CheckShopExists(userID uint64, shopName string) (bool, error) {
	conn := r.db.ConnectionGetting()

	count := int64(0)
	err := conn.Model(&entities.Shop{}).
		Where("user_id = ? AND name = ?", userID, shopName).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *shopRepositoryImpl) ListShops(userID uint64) ([]entities.Shop, error) {
	conn := r.db.ConnectionGetting()

	var shops []entities.Shop

	// Query shops for the specific user
	if err := conn.Where("user_id = ?", userID).Find(&shops).Error; err != nil {
		return nil, err
	}

	return shops, nil
}
