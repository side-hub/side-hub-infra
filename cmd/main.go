package main

import (
	"github.com/labstack/echo/v4"
	"github.com/side-hub-be/side-hub-be/internal/log"
	"github.com/side-hub-be/side-hub-be/internal/middleware"
	"github.com/side-hub-be/side-hub-be/internal/router"
)

const (
	Address = ":8080"
)

// @title		side-hub API
// @version		0.0.1
// @description	side hub backend api
// @BasePath	/
func main() {
	log.Init("")

	e := echo.New()
	middleware.Use(e)
	router.Route(e)

	log.Fatal(e.Start(Address))
}
