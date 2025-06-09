package service

import (
	"github.com/BoomTHDev/wear-pos-server/entities"
	"github.com/BoomTHDev/wear-pos-server/pkg/custom"
	_userModel "github.com/BoomTHDev/wear-pos-server/pkg/user/model"
)

type UserService interface {
	Add(user *entities.User) (*_userModel.RegisterResponse, *custom.AppError)
	List() ([]_userModel.UserResponse, *custom.AppError)
	Read(id uint64) (*entities.User, *custom.AppError)
}
