package repository

import "github.com/BoomTHDev/wear-pos-server/entities"

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
}
