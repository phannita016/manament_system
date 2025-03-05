package apps

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/phannita016/management/apis/serve"
)

// DefaultServe default struct collector.
// default port http-server-collector.
var DefaultServe = Server{
	Addr: ":8080",
	Echo: serve.New(),
}

type Server struct {
	Addr   string
	Echo   *echo.Echo
	Logger *slog.Logger
}

func NewServer(addr string, e *echo.Echo, logger *slog.Logger) *Server {
	return &Server{Addr: addr, Echo: e, Logger: logger}
}

// Serve function start-http-server collector-server.
func (s *Server) Run(ctx context.Context) error {
	s.Echo.HideBanner = true

	go func() {
		s.Logger.Debug("Server starting server at", slog.String("addr", s.Addr))
		if err := s.Echo.Start(s.Addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.Logger.Error("Server starting server at", slog.String("addr", s.Addr), slog.Any("err", err))
			s.Echo.Logger.Fatal(err)
		}
	}()

	s.Logger.Info("Server starting server at", slog.String("addr", s.Addr))
	return nil
}

// Shutdown stop-http-server collector-server.
func (s *Server) Stop(ctx context.Context) error {
	if err := s.Echo.Shutdown(ctx); err != nil {
		s.Logger.Error("Error shutting down server", slog.String("addr", s.Addr), slog.Any("err", err))
		return err
	}

	s.Logger.Info("Server shutting down", slog.String("addr", s.Addr))
	return nil
}
