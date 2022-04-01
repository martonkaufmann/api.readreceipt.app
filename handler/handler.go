package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/readreceipt/api/handler/public"
)

func RegisterRoutes(e *echo.Echo) {
	publicGroup := e.Group("/public")
	publicGroup.POST("/create", public.CreateHandler)
	publicGroup.GET("/read.png", public.ReadHandler)
	publicGroup.GET("/isread", public.IsReadHandler)
}
