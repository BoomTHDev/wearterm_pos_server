package server

import (
	_userController "github.com/BoomTHDev/wear-pos-server/pkg/user/controller"
	_userRepository "github.com/BoomTHDev/wear-pos-server/pkg/user/repository"
	_userService "github.com/BoomTHDev/wear-pos-server/pkg/user/service"
)

func (s *fiberServer) initUserRouter() {
	// Init User Plug
	userRepository := _userRepository.NewUserRepositoryImpl(s.db)
	userService := _userService.NewUserServiceImpl(userRepository)
	userController := _userController.NewUserController(userService)

	// User Routes
	userRouter := s.app.Group("/v1/users")
	userRouter.Get("/", userController.List)
	userRouter.Get("/:id", userController.Read)
}
