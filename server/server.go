package server

import (
	"time"

	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/readreceipt/api/server/request"
	"github.com/readreceipt/api/server/serializer"
)

func New() *echo.Echo {
	e := echo.New()

	e.JSONSerializer = serializer.JSON{}
	e.Validator = request.NewValidator()

	e.Server.ReadTimeout = time.Second * 5
	e.Server.WriteTimeout = time.Second * 15

	e.Use(middleware.BodyLimit("8K"))
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(sentryecho.New(sentryecho.Options{}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://readreceipt.localhost", "https://readreceipt.app"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	return e
}
