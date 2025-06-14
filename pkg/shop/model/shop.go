package model

import (
	"time"

	"github.com/BoomTHDev/wear-pos-server/entities"
)

type (
	NewShopRequest struct {
		Name   string `json:"name"`
		UserID uint64 `json:"user_id"`
	}

	ShopResponse struct {
		ID        uint64    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	}

	ShopResponses struct {
		ID        uint64    `json:"id"`
		Name      string    `json:"name"`
		UserID    uint64    `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
	}
)

func ToShopResponse(shop *entities.Shop) *ShopResponse {
	if shop == nil {
		return nil
	}
	return &ShopResponse{
		ID:   shop.ID,
		Name: shop.Name,
	}
}

func ToShopResponses(shops []entities.Shop) []ShopResponses {
	if shops == nil {
		return nil
	}

	listShops := []ShopResponses{}
	for _, shop := range shops {
		listShops = append(listShops, ShopResponses{
			ID:        shop.ID,
			Name:      shop.Name,
			UserID:    shop.UserID,
			CreatedAt: shop.CreatedAt,
		})
	}
	return listShops
}
