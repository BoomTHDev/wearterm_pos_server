package service

import (
	"github.com/BoomTHDev/wear-pos-server/entities"
	_userModel "github.com/BoomTHDev/wear-pos-server/pkg/user/model"
)

type UserService interface {
	Add(user *entities.User) (*_userModel.User, error)
}
