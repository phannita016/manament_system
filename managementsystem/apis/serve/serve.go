package serve

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(middlewareFunc ...echo.MiddlewareFunc) *echo.Echo {
	var e = echo.New()

	e.Use(middleware.RequestID())
	e.Use(middlewareFunc...)
	return e
}
