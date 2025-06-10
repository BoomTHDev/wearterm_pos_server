package service

import (
	"strconv"

	"github.com/BoomTHDev/wear-pos-server/config"
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

func (s *authServiceImpl) NewPIN(id uint64, pin int) *custom.AppError {
	// Check pin 000000-999999
	if pin <= 0 || pin > 999999 {
		return custom.ErrInvalidInput("INVALID_PIN", "Pin must be than 6 digits", nil)
	}

	// Check user exists
	existingUser, err := s.userRepository.ReadByID(id)
	if err != nil {
		if custom.IsRecordNotFoundError(err) {
			return custom.ErrNotFound("USER_NOT_FOUND", "User not found", err)
		}
		return custom.ErrIntervalServer("USER_READ_FAILED", "Failed to check user availability", err)
	}

	// Check pin already exists
	if existingUser.Pin != 0 {
		return custom.ErrConflict("USER_PIN_ALREADY_EXISTS", "User pin already exists", nil)
	}

	// Hash pin
	hashedPin, err := util.HashPassword(strconv.Itoa(pin))
	if err != nil {
		return custom.ErrIntervalServer("USER_CREATE_FAILED", "Failed to hash pin", err)
	}

	// Create pin
	if err := s.userRepository.CreatePIN(id, pin, hashedPin); err != nil {
		return custom.ErrIntervalServer("USER_CREATE_FAILED", "Failed to create user", err)
	}

	return nil
}

func (s *authServiceImpl) LoginWithPassword(req _userModel.LoginWithPasswordRequest) (*_userModel.LoginResponse, *custom.AppError) {
	// Check req
	if req.Username == "" || req.Password == "" {
		return nil, custom.ErrInvalidInput("INVALID_REQUEST_BODY", "Username or password is empty", nil)
	}

	// Check user existing
	existingUser, err := s.userRepository.ReadByUsername(req.Username)
	if err != nil {
		if custom.IsRecordNotFoundError(err) {
			return nil, custom.ErrNotFound("USER_NOT_FOUND", "User not found", err)
		}
		return nil, custom.ErrIntervalServer("USER_READ_FAILED", "Failed to check user availability", err)
	}

	// Check password
	if err := util.CheckPasswordHash(req.Password, existingUser.Password); err != nil {
		return nil, custom.ErrUnauthorized("INVALID_CREDENTIALS", "Invalid credentials", err)
	}

	// Generate token
	jwtSecret := config.ConfigGetting().Server.JWTSecret
	token, err := util.GenerateToken(existingUser.ID, jwtSecret)
	if err != nil {
		return nil, custom.ErrIntervalServer("TOKEN_GENERATION_FAILED", "Failed to generate token", err)
	}

	return &_userModel.LoginResponse{Token: token}, nil
}

func (s *authServiceImpl) LoginWithPin(req _userModel.LoginWithPinRequest) (*_userModel.LoginResponse, *custom.AppError) {
	// Check req
	if req.Username == "" || req.Pin == 0 {
		return nil, custom.ErrInvalidInput("INVALID_REQUEST_BODY", "Username or pin is empty", nil)
	}

	// Check user existing
	existingUser, err := s.userRepository.ReadByUsername(req.Username)
	if err != nil {
		if custom.IsRecordNotFoundError(err) {
			return nil, custom.ErrNotFound("USER_NOT_FOUND", "User not found", err)
		}
		return nil, custom.ErrIntervalServer("USER_READ_FAILED", "Failed to check user availability", err)
	}

	// Check user pin
	if existingUser.Pin == 0 {
		return nil, custom.ErrUnauthorized("INVALID_CREDENTIALS", "User pin is empty", nil)
	}

	// Check pin
	if err := util.CheckPasswordHash(strconv.Itoa(req.Pin), existingUser.HashPin); err != nil {
		return nil, custom.ErrUnauthorized("INVALID_CREDENTIALS", "Invalid credentials", err)
	}

	// Generate token
	jwtSecret := config.ConfigGetting().Server.JWTSecret
	token, err := util.GenerateToken(existingUser.ID, jwtSecret)
	if err != nil {
		return nil, custom.ErrIntervalServer("TOKEN_GENERATION_FAILED", "Failed to generate token", err)
	}

	return &_userModel.LoginResponse{Token: token}, nil
}
