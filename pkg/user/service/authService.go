package service

import (
	"github.com/BoomTHDev/wear-pos-server/pkg/custom"
	_userModel "github.com/BoomTHDev/wear-pos-server/pkg/user/model"
)

type AuthService interface {
	Register(req _userModel.RegisterRequest) (*_userModel.RegisterResponse, *custom.AppError)
	NewPIN(id uint64, pin int) *custom.AppError
	LoginWithPassword(req _userModel.LoginWithPasswordRequest) (*_userModel.LoginResponse, *custom.AppError)
	LoginWithPin(req _userModel.LoginWithPinRequest) (*_userModel.LoginResponse, *custom.AppError)
}
