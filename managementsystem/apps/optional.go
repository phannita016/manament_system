package apps

import (
	"github.com/labstack/echo/v4"
)

type Option func(*Apps)

func WithServeAddr(addr string) Option {
	return func(c *Apps) {
		c.Addr = addr
	}
}

func WithServeEcho(e *echo.Echo) Option {
	return func(c *Apps) {
		c.server = NewServer(c.Addr, e)
	}
}
