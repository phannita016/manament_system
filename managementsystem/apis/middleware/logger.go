package middleware

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	slogEcho "github.com/samber/slog-echo"
)

// Logger provider middleware control handler.
func Logger(logger *slog.Logger) echo.MiddlewareFunc {
	return slogEcho.New(logger.WithGroup("http"))
}
