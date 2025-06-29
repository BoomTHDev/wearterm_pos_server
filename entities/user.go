package entities

import (
	"time"
)

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Pin       int
	HashPin   string
	Shops     []Shop    `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}
