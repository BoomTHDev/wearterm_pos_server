package model

import "github.com/BoomTHDev/wear-pos-server/entities"

type (
	NewShopRequest struct {
		Name   string `json:"name"`
		UserID uint64 `json:"user_id"`
	}

	ShopResponse struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
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

func ToShopsResponse(shops []entities.Shop) []ShopResponse {
	if shops == nil {
		return nil
	}
	result := make([]ShopResponse, 0, len(shops))
	for _, shop := range shops {
		result = append(result, ShopResponse{
			ID:   shop.ID,
			Name: shop.Name,
		})
	}
	return result
}
