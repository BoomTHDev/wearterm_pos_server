package service

import (
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

func (s *userServiceImpl) List() ([]_userModel.UserResponses, *custom.AppError) {
	users, err := s.userRepository.List()
	if err != nil {
		return nil, custom.ErrIntervalServer("USER_LIST_FAILED", "Failed to list users", err)
	}
	return _userModel.ToUserResponses(users), nil
}

func (s *userServiceImpl) Read(id uint64) (*_userModel.UserResponse, *custom.AppError) {
	user, err := s.userRepository.ReadByID(id)
	if err != nil {
		if custom.IsRecordNotFoundError(err) {
			return nil, custom.ErrNotFound("USER_NOT_FOUND", "User not found", err)
		}
		return nil, custom.ErrIntervalServer("USER_READ_FAILED", "Failed to read user", err)
	}
	return _userModel.ToUserResponse(user), nil
}
