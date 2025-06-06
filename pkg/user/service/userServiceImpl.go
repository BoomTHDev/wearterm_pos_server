package service

import (
	"log"

	"github.com/BoomTHDev/wear-pos-server/entities"
	_userModel "github.com/BoomTHDev/wear-pos-server/pkg/user/model"
	_userRepository "github.com/BoomTHDev/wear-pos-server/pkg/user/repository"
	"github.com/gofiber/fiber/v2"
)

type userServiceImpl struct {
	userRepository _userRepository.UserRepository
}

func NewUserServiceImpl(userRepository _userRepository.UserRepository) UserService {
	return &userServiceImpl{userRepository: userRepository}
}

func (s *userServiceImpl) Add(user *entities.User) (*_userModel.User, error) {
	newUser, err := s.userRepository.Create(user)
	if err != nil {
		log.Printf("Add user error: %v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Interval server error")
	}
	return newUser.ToUserModel(), nil
}
