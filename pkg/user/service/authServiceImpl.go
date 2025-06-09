package service

import (
	"github.com/BoomTHDev/wear-pos-server/entities"
	"github.com/BoomTHDev/wear-pos-server/pkg/custom"
	_userModel "github.com/BoomTHDev/wear-pos-server/pkg/user/model"
	_userRepository "github.com/BoomTHDev/wear-pos-server/pkg/user/repository"
	"github.com/BoomTHDev/wear-pos-server/pkg/util"
)

type authServiceImpl struct {
	userRepository _userRepository.UserRepository
}

func NewAuthServiceImpl(userRepository _userRepository.UserRepository) AuthService {
	return &authServiceImpl{userRepository: userRepository}
}

func (s *authServiceImpl) Register(req _userModel.RegisterRequest) (*_userModel.RegisterResponse, *custom.AppError) {
	if req.Username == "" || req.Password == "" {
		return nil, custom.ErrInvalidInput("INVALID_REQUEST_BODY", "Username or password is empty", nil)
	}
	// Check if username already exists
	existingUser, err := s.userRepository.ReadByUsername(req.Username)
	if err == nil && existingUser != nil {
		return nil, custom.ErrConflict("USER_DUPLICATE", "User with this username already exists", nil)
	}
	// If error is not "not found" error, return the error
	if err != nil && !custom.IsRecordNotFoundError(err) {
		return nil, custom.ErrIntervalServer("USER_READ_FAILED", "Failed to check username availability", err)
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, custom.ErrIntervalServer("USER_CREATE_FAILED", "Failed to hash password", err)
	}

	user := entities.User{
		Username: req.Username,
		Password: hashedPassword,
	}
	newUser, err := s.userRepository.Create(&user)
	if err != nil {
		return nil, custom.ErrIntervalServer("USER_CREATE_FAILED", "Failed to create user", err)
	}

	return _userModel.ToRegisterResponse(newUser), nil
}
