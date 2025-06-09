package repository

import "github.com/BoomTHDev/wear-pos-server/entities"

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	List() ([]entities.User, error)
	Read(id uint64) (*entities.User, error)
}
