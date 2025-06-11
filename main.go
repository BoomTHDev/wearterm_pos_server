package main

import (
	"github.com/BoomTHDev/wear-pos-server/config"
	"github.com/BoomTHDev/wear-pos-server/databases"
	_ "github.com/BoomTHDev/wear-pos-server/docs"
	"github.com/BoomTHDev/wear-pos-server/server"
)

func main() {
	cfg := config.ConfigGetting()
	db := databases.NewPostgresDatabase(cfg.Database)
	server := server.NewFiberServer(cfg, db)

	server.Start()
}
