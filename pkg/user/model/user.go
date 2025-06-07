package model

import "github.com/BoomTHDev/wear-pos-server/entities"

type (
	RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	RegisterResponse struct {
		ID       uint64 `json:"id"`
		Username string `json:"username"`
	}

	ListUserResponse struct {
		*entities.User
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

func ToListUserResponse(users []entities.User) []ListUserResponse {
	if users == nil {
		return nil
	}
	listUsers := []ListUserResponse{}
	for _, user := range users {
		listUsers = append(listUsers, ListUserResponse{
			User: &user,
		})
	}
	return listUsers
}
