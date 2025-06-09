package model

import (
	"time"

	"github.com/BoomTHDev/wear-pos-server/entities"
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

	UserResponse struct {
		ID        uint64    `json:"id"`
		Username  string    `json:"username"`
		CreatedAt time.Time `json:"created_at"`
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

func ToUsersResponse(users []entities.User) []UserResponse {
	if users == nil {
		return nil
	}
	listUsers := []UserResponse{}
	for _, user := range users {
		listUsers = append(listUsers, UserResponse{
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
	}
}
