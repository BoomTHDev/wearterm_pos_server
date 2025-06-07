package entities

import (
	"time"

	_userModel "github.com/BoomTHDev/wear-pos-server/pkg/user/model"
)

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}

func (u *User) ToRegisterResponse() *_userModel.RegisterResponse {
	return &_userModel.RegisterResponse{
		ID:       u.ID,
		Username: u.Username,
	}
}
