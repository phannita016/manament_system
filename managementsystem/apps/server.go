package apps

import (
	"context"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/phannita016/management/apis/serve"
)

var DefaultServe = Server{
	Addr: ":8080",
	Echo: serve.New(),
}

type Server struct {
	Addr string
	Echo *echo.Echo
}

func NewServer(addr string, e *echo.Echo) *Server {
	return &Server{Addr: addr, Echo: e}
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		if err := s.Echo.Start(s.Addr); err != nil {
			slog.Error("Echo server failed to start", slog.Any("err", err))
			panic(err)
		}
	}()
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.Echo.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
