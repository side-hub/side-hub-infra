package router

import (
	"github.com/labstack/echo/v4"
	"github.com/side-hub-be/side-hub-be/internal/api"
)

// Route
func Route(e *echo.Echo) {
	e.GET("/health", api.GetHealth)
}
