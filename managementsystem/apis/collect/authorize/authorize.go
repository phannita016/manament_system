package authorize

import (
	"github.com/labstack/echo/v4"
	"github.com/phannita016/management/store"
)

type handleAuthorize struct {
	secret []byte

	cacheStore store.Cache[string]
}

func New(e *echo.Echo, secret []byte, cacheStore store.Cache[string]) {
	h := handleAuthorize{
		secret: secret,

		cacheStore: cacheStore,
	}

	g := e.Group("/authorize")
	g.POST("/login", h.Login)
	g.GET("/logout", h.Logout)
	g.POST("/refresh", h.RefreshToken)
}
