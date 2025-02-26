package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/phannita016/management/types"
)

func authorize(secret []byte, skippers []string) echojwt.Config {
	return echojwt.Config{
		Skipper: func(c echo.Context) bool {
			for _, skipper := range skippers {
				if strings.HasPrefix(c.Request().RequestURI, skipper) {
					return true
				}
			}
			return false
		},
		SigningMethod: jwt.SigningMethodHS256.Name,
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, err).WithInternal(err)
		},
		ContextKey: types.Authorization,
		SigningKey: secret,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(types.JwtCustomClaims)
		},
	}
}

// Authorization control access token.
func (m *Middleware) Authorization() echo.MiddlewareFunc {
	config := authorize(m.Secret, m.skippers)
	return echojwt.WithConfig(config)
}
