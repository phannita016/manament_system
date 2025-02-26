package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/phannita016/management/apps"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := apps.Apps{
		Addr:   ":8080",
		Ctx:    ctx,
		Secret: []byte("secret-key"),
	}

	opts := []apps.Option{}
	shutdown := apps.AppsServer(&app, opts...)

	if err := shutdown(ctx); err != nil {
		slog.Error("ServerShutdown", slog.Any("err", err))
	}
}
