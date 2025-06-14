package model

import (
	"time"

	"github.com/BoomTHDev/wear-pos-server/entities"
	_shopModel "github.com/BoomTHDev/wear-pos-server/pkg/shop/model"
)

type (
	RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	RegisterResponse struct {
		ID       uint64 `json:"id"`
		Username string `json:"username"`
	}

	UsersResponse struct {
		ID        uint64    `json:"id"`
		Username  string    `json:"username"`
		CreatedAt time.Time `json:"created_at"`
	}

	UserResponse struct {
		ID        uint64                    `json:"id"`
		Username  string                    `json:"username"`
		CreatedAt time.Time                 `json:"created_at"`
		Shops     []_shopModel.ShopResponse `json:"shops"`
	}

	NewPINRequest struct {
		Pin int `json:"pin"`
	}

	LoginWithPasswordRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginWithPinRequest struct {
		Username string `json:"username"`
		Pin      int    `json:"pin"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}
)

func ToRegisterResponse(user *entities.User) *RegisterResponse {
	if user == nil {
		return nil
	}
	return &RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}

func ToUsersResponse(users []entities.User) []UsersResponse {
	if users == nil {
		return nil
	}
	listUsers := []UsersResponse{}
	for _, user := range users {
		listUsers = append(listUsers, UsersResponse{
			ID:        user.ID,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
		})
	}
	return listUsers
}

func ToUserResponse(user *entities.User) *UserResponse {
	if user == nil {
		return nil
	}

	return &UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		Shops:     _shopModel.ToShopsResponse(user.Shops),
	}
}
