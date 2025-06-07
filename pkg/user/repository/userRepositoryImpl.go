package repository

import (
	"errors"

	"github.com/BoomTHDev/wear-pos-server/databases"
	"github.com/BoomTHDev/wear-pos-server/entities"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db databases.Database
}

func NewUserRepositoryImpl(db databases.Database) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Create(user *entities.User) (*entities.User, error) {
	conn := r.db.ConnectionGetting()

	// Create a new user record
	if err := conn.Create(user).Error; err != nil {
		return nil, err
	}

	// Fetch the created user to get all fields (including auto-generated ones)
	insertedUser := entities.User{}
	if err := conn.First(&insertedUser, user.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &insertedUser, nil
}

func (r *userRepositoryImpl) List() ([]entities.User, error) {
	conn := r.db.ConnectionGetting()

	users := []entities.User{}

	if err := conn.Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return users, nil
}
