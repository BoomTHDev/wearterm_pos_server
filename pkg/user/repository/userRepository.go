package repository

import "github.com/BoomTHDev/wear-pos-server/entities"

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	List() ([]entities.User, error)
	ReadByID(id uint64) (*entities.User, error)
	ReadByUsername(username string) (*entities.User, error)
	CreatePIN(id uint64, pin int, hashPin string) error
}
