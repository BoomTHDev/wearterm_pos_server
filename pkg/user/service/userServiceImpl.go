package service

import (
	"github.com/BoomTHDev/wear-pos-server/entities"
	"github.com/BoomTHDev/wear-pos-server/pkg/custom"
	_userModel "github.com/BoomTHDev/wear-pos-server/pkg/user/model"
	_userRepository "github.com/BoomTHDev/wear-pos-server/pkg/user/repository"
)

type userServiceImpl struct {
	userRepository _userRepository.UserRepository
}

func NewUserServiceImpl(userRepository _userRepository.UserRepository) UserService {
	return &userServiceImpl{userRepository: userRepository}
}

func (s *userServiceImpl) Add(user *entities.User) (*_userModel.User, *custom.AppError) {
	newUser, err := s.userRepository.Create(user)
	if err != nil {
		return nil, custom.ErrIntervalServer("INTERVAL_SERVER_ERROR", "Interval server error", err)
	}
	return newUser.ToUserModel(), nil
}
