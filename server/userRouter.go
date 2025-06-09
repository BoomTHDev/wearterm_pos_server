package server

import (
	_userController "github.com/BoomTHDev/wear-pos-server/pkg/user/controller"
	_userRepository "github.com/BoomTHDev/wear-pos-server/pkg/user/repository"
	_userService "github.com/BoomTHDev/wear-pos-server/pkg/user/service"
)

func (s *fiberServer) initUserRouter() {
	router := s.app.Group("/v1/users")

	userRepository := _userRepository.NewUserRepositoryImpl(s.db)
	userService := _userService.NewUserServiceImpl(userRepository)
	userController := _userController.NewUserController(userService)

	router.Post("/", userController.Add)
	router.Get("/", userController.List)
	router.Get("/:id", userController.Read)
}
