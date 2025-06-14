package server

import (
	_shopController "github.com/BoomTHDev/wear-pos-server/pkg/shop/controller"
	_shopRepository "github.com/BoomTHDev/wear-pos-server/pkg/shop/repository"
	_shopService "github.com/BoomTHDev/wear-pos-server/pkg/shop/service"
	"github.com/BoomTHDev/wear-pos-server/server/middleware"
)

func (s *fiberServer) initShopRouter() {
	// Init Shop Plug
	shopRepository := _shopRepository.NewShopRepositoryImpl(s.db)
	shopService := _shopService.NewShopServiceImpl(shopRepository)
	shopController := _shopController.NewShopControllerImpl(shopService)

	// Shop Routes
	shopRouter := s.app.Group("/v1/shop")
	shopRouter.Use(middleware.AuthMiddleware())
	shopRouter.Post("/:userId", shopController.NewShop)
	shopRouter.Get("/:userId", shopController.List)
}
