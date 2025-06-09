package repository

import (
	"github.com/BoomTHDev/wear-pos-server/databases"
	"github.com/BoomTHDev/wear-pos-server/entities"
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
		return nil, err
	}

	return &insertedUser, nil
}

func (r *userRepositoryImpl) List() ([]entities.User, error) {
	conn := r.db.ConnectionGetting()

	users := []entities.User{}

	if err := conn.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepositoryImpl) Read(id uint64) (*entities.User, error) {
	conn := r.db.ConnectionGetting()

	user := entities.User{}
	if err := conn.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
