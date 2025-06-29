package main

import (
	"github.com/BoomTHDev/wear-pos-server/config"
	"github.com/BoomTHDev/wear-pos-server/databases"
	"github.com/BoomTHDev/wear-pos-server/entities"
	"gorm.io/gorm"
)

func main() {
	cfg := config.ConfigGetting()
	db := databases.NewPostgresDatabase(cfg.Database)

	tx := db.ConnectionGetting().Begin()

	userMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func userMigration(tx *gorm.DB) {
	tx.AutoMigrate(&entities.User{}, &entities.Shop{})
}
