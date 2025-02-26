package middleware

import (
	"errors"
	"net/http"
	"slices"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/phannita016/management/store"
	"github.com/phannita016/management/types"
)

type Middleware struct {
	Secret   []byte
	cache    store.Cache[string]
	skippers []string
}

func NewMiddleware(secret []byte, cache store.Cache[string], skippers []string) *Middleware {
	return &Middleware{Secret: secret, cache: cache, skippers: skippers}
}

func (m *Middleware) SkipperURI(url string) bool {
	skip := []string{"/api/v1/authorize/logout"}

	idx := slices.Index(skip, url)
	return idx != -1
}

func (m *Middleware) AuthorizeWithToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if token, ok := types.Restricted(c); ok {
			if _, ok := m.cache.Get(token.ClaimID); !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorize").WithInternal(errors.New("token not found"))
			}

			if m.SkipperURI(c.Request().RequestURI) {
				return next(c)
			}

			if !token.IsAccess() {
				return echo.NewHTTPError(http.StatusNotAcceptable, "Authorize").WithInternal(jwt.ErrInvalidKey)
			}
		}

		return next(c)
	}
}
