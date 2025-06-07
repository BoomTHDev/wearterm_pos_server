package service

import (
	"github.com/BoomTHDev/wear-pos-server/entities"
	"github.com/BoomTHDev/wear-pos-server/pkg/custom"
	_userModel "github.com/BoomTHDev/wear-pos-server/pkg/user/model"
	_userRepository "github.com/BoomTHDev/wear-pos-server/pkg/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	userRepository _userRepository.UserRepository
}

func NewUserServiceImpl(userRepository _userRepository.UserRepository) UserService {
	return &userServiceImpl{userRepository: userRepository}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *userServiceImpl) Add(user *entities.User) (*_userModel.RegisterResponse, *custom.AppError) {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return nil, custom.ErrIntervalServer("USER_CREATE_FAILED", "Failed to hash password", err)
	}

	user.Password = hashedPassword
	newUser, err := s.userRepository.Create(user)
	if err != nil {
		return nil, custom.ErrIntervalServer("USER_CREATE_FAILED", "Failed to create user", err)
	}
	return _userModel.ToRegisterResponse(newUser), nil
}

func (s *userServiceImpl) List() ([]_userModel.ListUserResponse, *custom.AppError) {
	users, err := s.userRepository.List()
	if err != nil {
		return nil, custom.ErrIntervalServer("USER_LIST_FAILED", "Failed to list users", err)
	}
	return _userModel.ToListUserResponse(users), nil
}
