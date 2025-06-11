package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yokeTH/gofiber-scalar/scalar"
)

// ScalarHandler creates middleware for Scalar API Documentation
func ScalarHandler() fiber.Handler {
	// Note: Swagger registration is handled by the docs package
	return scalar.New(scalar.Config{
		BasePath: "/",
		Path:     "/scalar",
		Title:    "Wear POS Server API Documentation",
	})
}
