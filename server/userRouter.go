package server

import (
	_userController "github.com/BoomTHDev/wear-pos-server/pkg/user/controller"
	_userRepository "github.com/BoomTHDev/wear-pos-server/pkg/user/repository"
	_userService "github.com/BoomTHDev/wear-pos-server/pkg/user/service"
	"github.com/BoomTHDev/wear-pos-server/server/middleware"
)

func (s *fiberServer) initUserRouter() {
	// Init User Plug
	userRepository := _userRepository.NewUserRepositoryImpl(s.db)
	userService := _userService.NewUserServiceImpl(userRepository)
	userController := _userController.NewUserControllerImpl(userService)

	// User Routes
	userRouter := s.app.Group("/v1/users")
	userRouter.Use(middleware.AuthMiddleware())
	userRouter.Get("/", userController.List)
	userRouter.Get("/:id", userController.Read)
}
