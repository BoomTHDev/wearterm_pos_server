package server

import (
	_authController "github.com/BoomTHDev/wear-pos-server/pkg/user/controller"
	_userRepository "github.com/BoomTHDev/wear-pos-server/pkg/user/repository"
	_authService "github.com/BoomTHDev/wear-pos-server/pkg/user/service"
)

func (s *fiberServer) initAuthRouter() {
	// Init Auth Plug
	userRepository := _userRepository.NewUserRepositoryImpl(s.db)
	authService := _authService.NewAuthServiceImpl(userRepository)
	authController := _authController.NewAuthControllerImpl(authService)

	// Auth Routes
	authRouter := s.app.Group("/v1/auth")
	authRouter.Post("/register", authController.Register)
	authRouter.Put("/new-pin/:id", authController.NewPIN)
	authRouter.Post("/login", authController.LoginWithPassword)
	authRouter.Post("/login-with-pin", authController.LoginWithPin)
}
