package service

import (
	"github.com/BoomTHDev/wear-pos-server/pkg/custom"
	_userModel "github.com/BoomTHDev/wear-pos-server/pkg/user/model"
)

type UserService interface {
	List() ([]_userModel.UserResponse, *custom.AppError)
	Read(id uint64) (*_userModel.UserResponse, *custom.AppError)
}
