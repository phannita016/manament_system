package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/phannita016/management/apps"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	// a context that will be canceled
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := apps.Apps{
		Addr:   ":8080",
		Ctx:    ctx,
		Secret: []byte(os.Getenv("SECRET_KEY")),
	}
	logger := slog.Default()

	opts := []apps.Option{
		apps.WithLogger(logger),
	}
	shutdown := apps.AppsServer(&app, opts...)

	<-ctx.Done()

	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := shutdown(ctx); err != nil {
		slog.Error("ServerShutdown", slog.Any("err", err))
	}
}
