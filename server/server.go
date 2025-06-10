package server

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/BoomTHDev/wear-pos-server/config"
	"github.com/BoomTHDev/wear-pos-server/databases"
	"github.com/BoomTHDev/wear-pos-server/server/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

type fiberServer struct {
	app  *fiber.App
	db   databases.Database
	conf *config.Config
}

var (
	once   sync.Once
	server *fiberServer
)

func NewFiberServer(conf *config.Config, db databases.Database) *fiberServer {
	fiberApp := fiber.New(fiber.Config{
		BodyLimit:    conf.Server.BodyLimit * 1024 * 1024, // Convert MB to bytes
		IdleTimeout:  time.Second * time.Duration(conf.Server.TimeOut),
		ErrorHandler: middleware.ErrorHandler(),
	})

	once.Do(func() {
		server = &fiberServer{
			app:  fiberApp,
			db:   db,
			conf: conf,
		}
	})

	return server
}

func (s *fiberServer) Start() {
	s.app.Use(logger.New())

	// Join the allowed origins into a comma-separated string
	allowedOrigins := strings.Join(s.conf.Server.AllowOrigins, ",")

	s.app.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false, // Set to false when using wildcard
		// If you need to support credentials, you must specify exact origins
		// AllowCredentials: true,
		// Or use a function to dynamically check allowed origins
		// AllowOriginsFunc: func(origin string) bool {
		//     // Check if origin is in your allowed origins list
		//     for _, allowed := range s.conf.Server.AllowOrigins {
		//         if allowed == origin {
		//             return true
		//         }
		// }
		// return false
		// },
	}))

	s.app.Get("/swagger/*", swagger.HandlerDefault)
	// Health Check
	s.app.Get("/v1/health", s.healthCheck)

	s.initUserRouter()
	s.initAuthRouter()

	s.app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error": fiber.Map{
				"code":    "ROUTE_NOT_FOUND",
				"message": fmt.Sprintf("Sorry, endpoint %s %s not found", c.Method(), c.Path()),
			},
		})
	})

	s.httpListening()
}

func (s *fiberServer) httpListening() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Listen(url); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func (s *fiberServer) healthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Server is OK!",
	})
}
