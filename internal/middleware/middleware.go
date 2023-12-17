package middleware

import (
	"time"

	zeromw "github.com/faabiosr/echo-middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	zlog "github.com/rs/zerolog/log"
	echoSwagger "github.com/swaggo/echo-swagger"

	// swagger
	_ "github.com/side-hub-be/side-hub-be/docs"
)

const (
	RequestTimeout = 3 * time.Second
	JWTKey         = "secret"
)

func Use(e *echo.Echo) {
	// default
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middleware.ContextTimeout(RequestTimeout))

	// logger
	e.Use(zeroLog())

	// swagger
	e.GET("/swagger/:any", echoSwagger.WrapHandler)
}

func zeroLog() echo.MiddlewareFunc {
	zeromw.DefaultZeroLogConfig.Logger = zlog.Logger
	return zeromw.ZeroLogWithConfig(zeromw.DefaultZeroLogConfig)
}
