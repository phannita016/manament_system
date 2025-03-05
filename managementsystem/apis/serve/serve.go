package serve

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phannita016/management/apis/validate"
)

func New(middlewareFunc ...echo.MiddlewareFunc) *echo.Echo {
	var e = echo.New()

	e.Validator = validate.AddNew

	e.Use(middleware.RequestID()) // adds `X-Request-ID` header to the response
	e.Use(middleware.Logger())    // useful for debugging and monitoring API requests
	e.Use(middleware.Recover())   // prevents server crashes by handling panics
	e.Use(middleware.Secure())    // apply Secure Middleware to all routes
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	e.Use(middlewareFunc...)

	e.HTTPErrorHandler = HandleError
	return e
}
