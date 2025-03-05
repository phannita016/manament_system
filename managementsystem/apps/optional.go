package apps

import "log/slog"

type Option func(*Apps)

// WithLogger set echo pointer optional.
func WithLogger(logger *slog.Logger) Option {
	return func(c *Apps) {
		c.Logger = logger
	}
}

// func WithServeAddr(addr string) Option {
// 	return func(c *Apps) {
// 		c.Addr = addr
// 	}
// }

// func WithServeEcho(e *echo.Echo) Option {
// 	return func(c *Apps) {
// 		c.server = NewServer(c.Addr, e)
// 	}
// }
